package models

import "math/big"

type Token struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Symbol   string `json:"symbol,omitempty"`
	Decimals uint8  `json:"decimals,omitempty"`
}

type PoolSimple struct {
	PoolAddress string `json:"pool_address"`
	ChainName   string `json:"chain_name"`
	LPToken     Token  `json:"lp_token"`
}

type TokenBalance struct {
	TokenSymbol  string   `json:"token_symbol"`
	Balance      *big.Int `json:"balance"` // Should change to type of string to avoid number overflow?
	BalanceInUSD float64  `json:"balance_in_usd"`
}
