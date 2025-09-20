package repo

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewProtocolPositionRepository,
	NewUserTokenRepository,
	NewWalletAddressRepository,
	NewWalletAssetSnapshotRepository,
	NewTransactionRepository,
)
