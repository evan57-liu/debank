package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/coin50etf/coin-market/internal/model"
	"github.com/coin50etf/coin-market/internal/pkg/database"
	"github.com/coin50etf/coin-market/internal/pkg/logger"
	"github.com/coin50etf/coin-market/internal/pkg/third_party/debank"
	"github.com/coin50etf/coin-market/internal/pkg/third_party/debanksign"
	"github.com/coin50etf/coin-market/internal/repo"
)

type TransactionService struct {
	protocolPositionRepo *repo.ProtocolPositionRepository
	walletAddressRepo    *repo.WalletAddressRepository
	transactionRepo      *repo.TransactionRepository
	debankClient         *debank.Client
	debankSignClient     *debanksign.Client
	postgresDB           *database.PostgresDB
}

func NewTransactionService(
	protocolPositionRepo *repo.ProtocolPositionRepository,
	walletAddressRepo *repo.WalletAddressRepository,
	transactionRepo *repo.TransactionRepository,
	debankClient *debank.Client,
	debankSignClient *debanksign.Client,
	postgresDB *database.PostgresDB,
) *TransactionService {
	return &TransactionService{
		protocolPositionRepo: protocolPositionRepo,
		walletAddressRepo:    walletAddressRepo,
		transactionRepo:      transactionRepo,
		debankClient:         debankClient,
		debankSignClient:     debankSignClient,
		postgresDB:           postgresDB,
	}
}

func (p *TransactionService) ProcessTransaction(ctx context.Context) error {
	addresses, err := p.walletAddressRepo.FindAll()
	if err != nil {
		return fmt.Errorf("FindAll WalletAddress failed: %w", err)
	}
	syncTime := time.Now()

	transactions := make([]*model.Transaction, 0)
	for _, address := range addresses {
		debankSignDto, err := p.debankSignClient.GetSignature(ctx, address.Address)
		if err != nil {
			logger.Error(ctx, "GetSignature failed", "error", err, "address", address)
			continue
		}
		logger.Info(ctx, "get signature", "debank_sign", debankSignDto)

		debankResponse, err := p.debankClient.GetAllTransactions(ctx, address.Address, debankSignDto)
		if err != nil {
			logger.Error(ctx, "GetAllTransactions failed", "error", err, "address", address)
			continue
		}

		debankResultDto := debankResponse.Data.Result
		if debankResultDto == nil {
			logger.Error(ctx, "GetAllTransactions result is nil", "address", address)
			continue
		}

		for _, history := range debankResultDto.Data.HistoryList {
			if history.IsScam {
				continue
			}
			jsonBytes, err := json.Marshal(history)
			if err != nil {
				logger.Error(ctx, "Marshal history failed", "error", err, "history", history)
				continue
			}

			transaction := &model.Transaction{
				WalletAddress: address.Address,
				ChainID:       history.Chain,
				TxHash:        history.ID,
				Name:          history.Tx.Name,
				FromAddress:   history.Tx.FromAddr,
				ToAddress:     history.Tx.ToAddr,
				Detail:        string(jsonBytes),
				SyncTime:      syncTime,
			}
			transactions = append(transactions, transaction)
		}
	}

	if len(transactions) > 0 {
		if err := p.postgresDB.DB.Transaction(func(tx *gorm.DB) error {
			for _, address := range addresses {
				if err := p.transactionRepo.DeleteByCondition(map[string]interface{}{
					"wallet_address": address.Address,
				}, tx); err != nil {
					logger.Error(ctx, "DeleteByCondition Transaction failed", "error", err, "address", address)
					return fmt.Errorf("DeleteByCondition Transaction failed: %w", err)
				}
			}

			if err = p.transactionRepo.CreateInBatches(transactions, 100, tx); err != nil {
				logger.Error(ctx, "BatchUpsert Transaction failed", "error", err)
				return fmt.Errorf("BatchUpsert Transaction failed: %w", err)
			}

			return nil
		}); err != nil {
			return err
		}

	}

	return nil
}
