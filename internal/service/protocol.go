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
	protocolPositionRepo    *repo.ProtocolPositionRepository
	userTokenRepo           *repo.UserTokenRepository
	walletAddressRepo       *repo.WalletAddressRepository
	walletAssetSnapshotRepo *repo.WalletAssetSnapshotRepository
	debankClient            *debank.Client
	postgresDB              *database.PostgresDB
}

func NewProtocolService(
	protocolPositionRepo *repo.ProtocolPositionRepository,
	userTokenRepo *repo.UserTokenRepository,
	walletAddressRepo *repo.WalletAddressRepository,
	walletAssetSnapshotRepo *repo.WalletAssetSnapshotRepository,
	debankClient *debank.Client,
	postgresDB *database.PostgresDB,
) *ProtocolService {
	return &ProtocolService{
		protocolPositionRepo:    protocolPositionRepo,
		userTokenRepo:           userTokenRepo,
		walletAddressRepo:       walletAddressRepo,
		walletAssetSnapshotRepo: walletAssetSnapshotRepo,
		debankClient:            debankClient,
		postgresDB:              postgresDB,
	}
}

func (p *ProtocolService) ProcessProtocol(ctx context.Context) error {
	addresses, err := p.walletAddressRepo.FindAll()
	if err != nil {
		return fmt.Errorf("FindAll WalletAddress failed: %w", err)
	}
	syncTime := time.Now()

	protocolPositions := make([]*model.ProtocolPosition, 0)
	for _, address := range addresses {
		protocols, err := p.debankClient.GetUserAllSimpleProtocolList(ctx, address.Address)
		if err != nil {
			logger.Error(ctx, "GetUserAllSimpleProtocolList failed", "error", err, "address", address)
			continue
		}

		for _, simpleProtocol := range protocols {
			protocol, err := p.debankClient.GetUserProtocol(ctx, address.Address, simpleProtocol.ID)
			if err != nil {
				logger.Error(ctx, "GetUserProtocol failed", "error", err, "address", address, "protocolID", protocol.ID)
				continue
			}

			poolMap := make(map[string][]*dto.PortfolioItemDto)
			for _, portfolioItem := range protocol.PortfolioItemList {
				if _, exists := poolMap[portfolioItem.Pool.ID]; !exists {
					poolMap[portfolioItem.Pool.ID] = make([]*dto.PortfolioItemDto, 0)
				}
				poolMap[portfolioItem.Pool.ID] = append(poolMap[portfolioItem.Pool.ID], portfolioItem)
			}

			poolName := ""
			description := ""
			for poolID, portfolioItems := range poolMap {
				assetTokens := make([]*dto.TokenDto, 0)
				for i, portfolioItem := range portfolioItems {
					/*if strings.ToLower(portfolioItem.Name) == "vesting" ||
						strings.ToLower(portfolioItem.Name) == "locked" {
						continue
					}*/

					for _, token := range portfolioItem.AssetTokenList {
						token.Type = "balance"
						if portfolioItem.Detail.RewardTokenList != nil {
							for _, rewardToken := range *portfolioItem.Detail.RewardTokenList {
								if rewardToken.ID == token.ID {
									token.Type = "reward"
									break
								}
							}
						}

						if i == 0 && token.Type == "balance" {
							description += token.Symbol + "+"
						}
					}

					if i == 0 {
						poolName = portfolioItem.Name
						if len(description) > 0 {
							description = description[:len(description)-1]
						}
						if len(portfolioItem.Detail.Description) > 0 {
							description = portfolioItem.Detail.Description
						}
					}

					assetTokens = append(assetTokens, portfolioItem.AssetTokenList...)
				}

				assetTokensJson, err := json.Marshal(assetTokens)
				if err != nil {
					logger.Error(ctx, "Marshal assetTokens failed", "error", err)
					continue
				}

				protocolPosition := &model.ProtocolPosition{
					Address:         address.Address,
					ChainID:         protocol.Chain,
					ProtocolID:      protocol.ID,
					ProtocolName:    protocol.Name,
					PoolID:          poolID,
					PoolName:        poolName,
					PoolDescription: description,
					AssetTokens:     string(assetTokensJson),
					SyncTime:        syncTime,
				}
				protocolPositions = append(protocolPositions, protocolPosition)
			}
		}
	}

	if len(protocolPositions) > 0 {
		if err = p.protocolPositionRepo.CreateInBatches(protocolPositions, 100, p.postgresDB.DB); err != nil {
			logger.Error(ctx, "BatchUpsert ProtocolPosition failed", "error", err)
			return fmt.Errorf("BatchUpsert ProtocolPosition failed: %w", err)
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
