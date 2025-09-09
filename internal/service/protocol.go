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
	protocolMappingRepo  *repo.ProtocolMappingRepository
	protocolPositionRepo *repo.ProtocolPositionRepository
	userTokenRepo        *repo.UserTokenRepository
	walletAddressRepo    *repo.WalletAddressRepository
	debankClient         *debank.Client
	postgresDB           *database.PostgresDB
}

func NewProtocolService(
	protocolMappingRepo *repo.ProtocolMappingRepository,
	protocolPositionRepo *repo.ProtocolPositionRepository,
	userTokenRepo *repo.UserTokenRepository,
	walletAddressRepo *repo.WalletAddressRepository,
	debankClient *debank.Client,
	postgresDB *database.PostgresDB,
) *ProtocolService {
	return &ProtocolService{
		protocolMappingRepo:  protocolMappingRepo,
		protocolPositionRepo: protocolPositionRepo,
		userTokenRepo:        userTokenRepo,
		walletAddressRepo:    walletAddressRepo,
		debankClient:         debankClient,
		postgresDB:           postgresDB,
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
				assetTokens = append(assetTokens, item.AssetTokenList...)
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

	return nil
}
