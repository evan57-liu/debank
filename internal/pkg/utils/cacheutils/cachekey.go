package cacheutils

const (
	NamespaceSymbol    = "symbol"
	NamespaceKLine     = "k_line"
	NamespaceMexc      = "mexc"
	NamespaceBinance   = "binance"
	NamespaceCcPayment = "cc_payment"
	NamespaceEtfKLiine = "etf_k_line"
	NamespaceEtfSymbol = "etf_symbol"
)

func GetMexcRateLimitKey() string {
	return "mexc_rate_limit"
}

func GetBinaceRateLimitKey() string {
	return "binance_rate_limit"
}

func GetSymbolsKey(any any) string {
	return "symbols:" + generateCacheKey(any)
}

func GetSymbolKLinesKey(any any) string {
	return "symbol_k_lines:" + generateCacheKey(any)
}

func GetMultiSymbolKLinesKey(any any) string {
	return "multi_symbol_k_lines:" + generateCacheKey(any)
}

func GetFixedSortSymbolsKey(any any) string {
	return "fixed_sort_symbols" + generateCacheKey(any)
}

func GetSymbolRanksKey(any any) string {
	return "symbol_ranks:" + generateCacheKey(any)
}

func GetCoinsKey() string {
	return "coins"
}

func GetEtfTrendsKey(any any) string {
	return "etf_k_line_trends:" + generateCacheKey(any)
}

func GetSymbolDailyKLinesKey(any any) string {
	return "symbol_daily_k_lines:" + generateCacheKey(any)
}

func GetEtfHoldingsKey(any any) string {
	return "etf_holdings:" + generateCacheKey(any)
}
