package main

import "fmt"

func main() {
	// ETF 初始配置
	etf := ETFPortfolio{
		//TotalShares: 1, // ETF 份额 S_0
		Tokens: []Token{
			{Name: "BTC", Percentage: 0.5, InitialPrice: 50000}, // 50%
			{Name: "ETH", Percentage: 0.3, InitialPrice: 3000},  // 30%
			{Name: "USDT", Percentage: 0.2, InitialPrice: 1},    // 20%
		},
	}

	// 模拟获取交易所 API 实时价格
	latestPrices := map[string]float64{
		"BTC":  100000, // BTC 价格更新
		"ETH":  3000,   // ETH 价格更新
		"USDT": 1,      // USDT 不变
	}

	// 计算最新净值
	newNav := etf.GetNewNav(latestPrices)
	fmt.Printf("最新的ETF净值: %.2f USDT/份\n", newNav)
}

type Token struct {
	Name         string  // 币种名称
	Percentage   float64 // 币种占比
	InitialPrice float64 // 币种初始价格
}

type ETFPortfolio struct {
	TotalShares float64 // 总 ETF 份额 (S)
	Tokens      []Token // 币种组成
}

func (etf *ETFPortfolio) GetNewNav(latestPrices map[string]float64) float64 {
	initialNav := 100.0

	// 计算币种的持仓数量
	holdings := make(map[string]float64)
	for _, token := range etf.Tokens {
		if token.InitialPrice == 0 {
			return initialNav
		}
		holdings[token.Name] = (token.Percentage * initialNav) / token.InitialPrice
	}

	// 计算新的总资产价值
	var newValue float64
	for name, qty := range holdings {
		fmt.Println(name, qty)
		newPrice, exists := latestPrices[name]
		if !exists {
			fmt.Printf("错误: 缺少 %s 的实时市场价格\n", name)
			return 0
		}
		newValue += qty * newPrice
	}
	fmt.Println("==========", newValue)

	// 计算新的 NAV
	return newValue
}
