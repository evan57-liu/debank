package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/coin50etf/coin-market/internal/dto"
	"github.com/coin50etf/coin-market/internal/model"
	"github.com/coin50etf/coin-market/internal/pkg/database"
	"github.com/coin50etf/coin-market/internal/pkg/logger"
	"github.com/coin50etf/coin-market/internal/pkg/third_party/debank"
	"github.com/coin50etf/coin-market/internal/repo"
)

type ProtocolService struct {
	protocolMappingRepo     *repo.ProtocolMappingRepository
	protocolPositionRepo    *repo.ProtocolPositionRepository
	userTokenRepo           *repo.UserTokenRepository
	walletAddressRepo       *repo.WalletAddressRepository
	walletAssetSnapshotRepo *repo.WalletAssetSnapshotRepository
	debankClient            *debank.Client
	postgresDB              *database.PostgresDB
}

func NewProtocolService(
	protocolMappingRepo *repo.ProtocolMappingRepository,
	protocolPositionRepo *repo.ProtocolPositionRepository,
	userTokenRepo *repo.UserTokenRepository,
	walletAddressRepo *repo.WalletAddressRepository,
	walletAssetSnapshotRepo *repo.WalletAssetSnapshotRepository,
	debankClient *debank.Client,
	postgresDB *database.PostgresDB,
) *ProtocolService {
	return &ProtocolService{
		protocolMappingRepo:     protocolMappingRepo,
		protocolPositionRepo:    protocolPositionRepo,
		userTokenRepo:           userTokenRepo,
		walletAddressRepo:       walletAddressRepo,
		walletAssetSnapshotRepo: walletAssetSnapshotRepo,
		debankClient:            debankClient,
		postgresDB:              postgresDB,
	}
}

func (p *ProtocolService) ProcessProtocol(ctx context.Context) error {
	protocolMappings, err := p.protocolMappingRepo.FindAll()
	if err != nil {
		return fmt.Errorf("FindAll ProtocolMapping failed: %w", err)
	}

	syncTime := time.Now()
	protocolPositions := make([]*model.ProtocolPosition, 0)
	for _, mapping := range protocolMappings {
		protocolPools := mapping.ProtocolPools

		for _, pool := range protocolPools {
			if pool.Status == 1 {
				continue
			}

			assetTokens := make([]*dto.TokenDto, 0)
			protocol, err := p.debankClient.GetUserProtocol(ctx, pool.Address, pool.ProtocolID)
			if err != nil {
				logger.Error(ctx, "GetProtocolPositionList failed", "error", err)
				continue
			}
			for _, item := range protocol.PortfolioItemList {
				if item.Pool.ID != pool.PoolID {
					continue
				}
				for _, token := range item.AssetTokenList {
					token.Type = "balance"
					if item.Detail.RewardTokenList != nil {
						for _, rewardToken := range *item.Detail.RewardTokenList {
							if rewardToken.ID == token.ID {
								token.Type = "reward"
								break
							}
						}
					}
				}
				assetTokens = append(assetTokens, item.AssetTokenList...)
			}

			if len(assetTokens) == 0 {
				pool.Status = 1
			}
			assetTokensJson, err := json.Marshal(assetTokens)
			if err != nil {
				logger.Error(ctx, "Marshal assetTokens failed", "error", err)
				continue
			}

			protocolPosition := &model.ProtocolPosition{
				InternalProtocolName: mapping.InternalProtocolName,
				InternalProtocolID:   mapping.InternalProtocolID,
				Address:              pool.Address,
				ChainID:              pool.ChainID,
				ProtocolID:           pool.ProtocolID,
				PoolID:               pool.PoolID,
				CustomID:             fmt.Sprintf("%s.%s.%s.%s", pool.Address, pool.ChainID, pool.ProtocolID, pool.PoolID),
				AssetTokens:          string(assetTokensJson),
				SyncTime:             syncTime,
			}
			protocolPositions = append(protocolPositions, protocolPosition)
		}
	}

	if len(protocolPositions) > 0 {
		err = p.protocolPositionRepo.CreateInBatches(protocolPositions, 100, p.postgresDB.DB)
		if err != nil {
			logger.Error(ctx, "BatchUpsert ProtocolPosition failed", "error", err)
			return fmt.Errorf("BatchUpsert ProtocolPosition failed: %w", err)
		}
	}

	for _, mapping := range protocolMappings {
		jsonBytes, err := json.Marshal(mapping.ProtocolPools)
		if err != nil {
			logger.Error(ctx, "Marshal ProtocolPools failed", "error", err)
			continue
		}

		if err := p.protocolMappingRepo.UpdateByCondition(map[string]interface{}{
			"id": mapping.ID,
		}, map[string]interface{}{
			"protocol_pools": string(jsonBytes),
		}); err != nil {
			logger.Error(ctx, "failed to update system country", "err", err)
		}
	}

	return nil
}

func (p *ProtocolService) ProcessUserTokens(ctx context.Context) error {
	addresses, err := p.walletAddressRepo.FindAll()
	if err != nil {
		return fmt.Errorf("FindAll WalletAddress failed: %w", err)
	}

	syncTime := time.Now()

	for _, address := range addresses {
		chains, err := p.debankClient.GetUserChainList(ctx, address.Address)
		if err != nil {
			return fmt.Errorf("GetUserChainList failed: %w", err)
		}

		userTokens := make([]*model.UserToken, 0)
		for _, chain := range chains {
			tokens, err := p.debankClient.GetUserTokenList(ctx, address.Address, chain.Id)
			if err != nil {
				logger.Error(ctx, "GetUserTokenList failed", "error", err, "address", address, "chain", chain.Id)
				continue
			}
			for _, token := range tokens {
				if !token.IsWallet {
					continue
				}
				userToken := &model.UserToken{
					Address:        address.Address,
					ChainID:        chain.Id,
					ContractID:     token.Id,
					Symbol:         token.Symbol,
					Decimals:       token.Decimals,
					LogoUrl:        token.LogoUrl,
					Price:          token.Price,
					Price24HChange: token.Price24HChange,
					TimeAt:         token.TimeAt,
					Amount:         token.Amount,
					SyncTime:       syncTime,
				}

				userTokens = append(userTokens, userToken)
			}
		}

		if len(userTokens) > 0 {
			err = p.userTokenRepo.CreateInBatches(userTokens, 100, p.postgresDB.DB)
			if err != nil {
				return fmt.Errorf("BatchUpsert UserToken failed: %w", err)
			}
		}
	}

	walletAssets := make([]*model.WalletAssetSnapshot, 0)
	for _, address := range addresses {
		balance, err := p.debankClient.GetUserTotalBalance(ctx, address.Address)
		if err != nil {
			logger.Error(ctx, "GetUserTotalBalance failed", "error", err, "address", address)
			continue
		}

		walletAssets = append(walletAssets, &model.WalletAssetSnapshot{
			SyncTime:      syncTime,
			WalletAddress: address.Address,
			TotalUSDValue: balance.TotalUsdValue,
		})
	}

	if len(walletAssets) > 0 {
		err = p.walletAssetSnapshotRepo.CreateInBatches(walletAssets, 100, p.postgresDB.DB)
		if err != nil {
			return fmt.Errorf("BatchUpsert WalletAssetSnapshot failed: %w", err)
		}
	}

	return nil
}
