package statistics_v2

type Summary struct {
	// Profit is the sum of all profits in points
	Profit int `json:"profit"`
	// AverageProfit is the average profit rounded to the nearest integer
	// I chose to round to the nearest integer because at the end this is result in points, not pips
	AverageProfit int `json:"averageProfit"`
	// WinLossRatio is the ratio of winning trades to losing trades
	// warning! break even trades are not taken into account
	WinLossRatio float64 `json:"winLossRatio"`
	// BestTrade is the Trade with the highest profit
	BestTrade Trade `json:"bestTrade"`
	// WorstTrade is the Trade with the lowest profit
	WorstTrade Trade `json:"worstTrade"`
	// BySymbol shows statistics for each symbol
	BySymbol map[string]BySymbol `json:"bySymbol"`
}

type BySymbol struct {
	// Profit is the sum of all profits in points
	Profit int `json:"profit"`
	// AverageProfit is the average profit rounded to the nearest integer
	// I chose to round to the nearest integer because at the end this is result in points, not pips
	AverageProfit int `json:"averageProfit"`
	// Amount is the number of trades
	Amount int `json:"amount"`
	// PercentOfAll is the percentage of all trades
	PercentOfAll int `json:"percentOfAll"`
}
