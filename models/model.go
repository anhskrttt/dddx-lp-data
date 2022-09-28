package models

type UserInformation struct {
	Address        string  `json:"address"`
	Protocol       string  `json:"protocol"`
	Tag            string  `json:"tag"`
	PairID         string  `json:"pair_id"`
	PairSymbol     string  `json:"pair_symbol"`
	PairName       string  `json:"pair_name"`
	Token0Bal      int     `json:"token0_bal"`
	Token1Bal      int     `json:"token1_bal"`
	Token0BalInUsd float64 `json:"token0_bal_in_usd"`
	Token1BalInUsd float64 `json:"token1_bal_in_usd"`
}
