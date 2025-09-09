package dto

import "github.com/shopspring/decimal"

type ProtocolDto struct {
	ID                    string              `json:"id"`
	Chain                 string              `json:"chain"`
	Name                  string              `json:"name"`
	SiteURL               string              `json:"site_url"`
	LogoURL               string              `json:"logo_url"`
	HasSupportedPortfolio bool                `json:"has_supported_portfolio"`
	TVL                   float64             `json:"tvl"`
	PortfolioItemList     []*PortfolioItemDto `json:"portfolio_item_list"`
}

type PortfolioItemDto struct {
	Stats          StatsDto               `json:"stats"`
	AssetDict      map[string]float64     `json:"asset_dict"`
	AssetTokenList []*TokenDto            `json:"asset_token_list"`
	UpdateAt       float64                `json:"update_at"`
	Name           string                 `json:"name"`
	DetailTypes    []string               `json:"detail_types"`
	Detail         PortfolioDetailDto     `json:"detail"`
	ProxyDetail    map[string]interface{} `json:"proxy_detail"` // 可根据实际扩展
	Pool           PoolDto                `json:"pool"`
}

type StatsDto struct {
	AssetUSDValue float64 `json:"asset_usd_value"`
	DebtUSDValue  float64 `json:"debt_usd_value"`
	NetUSDValue   float64 `json:"net_usd_value"`
}

type TokenDto struct {
	ID              string  `json:"id"`
	Chain           string  `json:"chain"`
	Name            string  `json:"name"`
	Symbol          string  `json:"symbol"`
	DisplaySymbol   *string `json:"display_symbol"`
	OptimizedSymbol string  `json:"optimized_symbol"`
	Decimals        int     `json:"decimals"`
	LogoURL         string  `json:"logo_url"`
	ProtocolID      string  `json:"protocol_id"`
	Price           float64 `json:"price"`
	Price24HChange  float64 `json:"price_24h_change"`
	CreditScore     float64 `json:"credit_score"`
	IsVerified      bool    `json:"is_verified"`
	IsScam          bool    `json:"is_scam"`
	IsSuspicious    *bool   `json:"is_suspicious"`
	IsCore          bool    `json:"is_core"`
	TotalSupply     float64 `json:"total_supply"`
	IsWallet        bool    `json:"is_wallet"`
	TimeAt          float64 `json:"time_at"`
	LowCreditScore  bool    `json:"low_credit_score"`
	Amount          float64 `json:"amount"`
}

type PortfolioDetailDto struct {
	SupplyTokenList *[]TokenDto `json:"supply_token_list"`
	RewardTokenList *[]TokenDto `json:"reward_token_list,omitempty"`
	BorrowTokenList *[]TokenDto `json:"borrow_token_list,omitempty"`
	Description     string      `json:"description,omitempty"`
}

type PoolDto struct {
	ID         string `json:"id"`
	Chain      string `json:"chain"`
	ProjectID  string `json:"project_id"`
	AdapterID  string `json:"adapter_id"`
	Controller string `json:"controller"`
	Index      *int   `json:"index"`
	TimeAt     int64  `json:"time_at"`
}

type UserTokenDto struct {
	Id              string           `json:"id"`
	Chain           string           `json:"chain"`
	Name            string           `json:"name"`
	Symbol          string           `json:"symbol"`
	DisplaySymbol   *string          `json:"display_symbol"`
	OptimizedSymbol string           `json:"optimized_symbol"`
	Decimals        int              `json:"decimals"`
	LogoUrl         *string          `json:"logo_url"`
	ProtocolId      string           `json:"protocol_id"`
	Price           decimal.Decimal  `json:"price"`
	Price24HChange  *decimal.Decimal `json:"price_24h_change"`
	IsVerified      bool             `json:"is_verified"`
	IsCore          bool             `json:"is_core"`
	IsWallet        bool             `json:"is_wallet"`
	TimeAt          float64          `json:"time_at"`
	TotalSupply     decimal.Decimal  `json:"total_supply"`
	CreditScore     decimal.Decimal  `json:"credit_score"`
	Amount          decimal.Decimal  `json:"amount"`
}
type UserUsedChainDto struct {
	BornAt           int    `json:"born_at"`
	Id               string `json:"id"`
	CommunityId      int    `json:"community_id"`
	Name             string `json:"name"`
	NativeTokenId    string `json:"native_token_id"`
	LogoUrl          string `json:"logo_url"`
	WrappedTokenId   string `json:"wrapped_token_id"`
	IsSupportPreExec bool   `json:"is_support_pre_exec"`
}
