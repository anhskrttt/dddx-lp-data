package models

import "math/big"

type User struct {
	Address   string     `json:"address"`
	Protocols []Protocol `json:"protocols"`
}

type Token struct {
	ID       string `json:"id,,omitempty"`
	Name     string `json:"name,omitempty"`
	Symbol   string `json:"symbol,omitempty"`
	Decimals uint8  `json:"decimals,omitempty"`
}

type PoolSimple struct {
	PoolAddress string `json:"pool_address"`
	ChainName   string `json:"chain_name"`
	LPToken     Token  `json:"lp_token"`
}

type Pool struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Symbol      string   `json:"symbol"`
	Address     string   `json:"address"`
	Chain       Chain    `json:"chain"`
	URL         string   `json:"url"`
	Description string   `json:"description"`
	Gecko_ID    string   `json:"gecko_id"`
	Tags        []string // E.g. DEX, DEX-Aggregator, Lending, etc.
	TVL         float64
}

type ProtocolSimple struct {
	Name    string `json:"name"`
	Address string `json:"string"`
	Chain   string `json:"chain"`
}

type Protocol struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Symbol      string   `json:"symbol"`
	Address     string   `json:"address"`
	Chain       Chain    `json:"chain"`
	URL         string   `json:"url"`
	Description string   `json:"description"`
	Gecko_ID    string   `json:"gecko_id"`
	Tags        []string // E.g. DEX, DEX-Aggregator, Lending, etc.
	TVL         float64
}

type Chain struct {
	/****************************/
	/* General information */
	ID          string `json:"id"`
	Name        string `json:"name"`
	NativeToken Token  `json:"native_token"`
	URL         string `json:"url"`
	Explorer    string `json:"explorer"`
	/* End of general information */
	/****************************/

	/****************************/
	/* Development information */
	RPC       []string `json:"rpc"`
	Faucet    []string `json:"faucet"`
	ChainId   string   `json:"chain_id"`
	NetworkId string   `json:"network_id"`
	// DefiProtocols []Protocol `json:"defi_protocol"`

	/* End of development information */
	/****************************/

}

type TokenBalance struct {
	TokenSymbol  string   `json:"token_symbol"`
	Balance      *big.Int `json:"balance"` // Should change to type of string to avoid number overflow?
	BalanceInUSD float64  `json:"balance_in_usd"`
}

type TokenPairBalance struct {
	ID        string       `json:"id"`
	Token0Bal TokenBalance `json:"token0_balance"`
	Token1Bal TokenBalance `json:"token1_balance"`
}

type UserBalance struct {
	User      User             `json:"user"`
	Type      BalanceType      `json:"type"`
	TokenPair TokenPairBalance `json:"token_pair"`
}

type Protocol_Balance struct {
	ID             string             `json:"id"`
	UserAddress    string             `json:"user"`
	Protocol       Protocol           `json:"protocol"`
	LPBalance      []TokenPairBalance `json:"lp_balance"`
	FarmingBalance []TokenPairBalance `json:"farming_balance"`
	StakedBalance  []TokenBalance     `json:"staked_balance"`
	LendingBalance []TokenBalance     `json:"lending_balance"`
	BorrowBalance  []TokenBalance     `json:"borrow_balance"`
}

// Utils
// Old struct

type UserInformation struct {
	Address        string  `json:"address"`
	Protocol       string  `json:"protocol"`
	Tag            string  `json:"tag"`
	PairID         string  `json:"pair_id"`
	PairSymbol     string  `json:"pair_symbol"`
	PairName       string  `json:"pair_name"`
	Token0Bal      string  `json:"token0_bal"`
	Token1Bal      string  `json:"token1_bal"`
	Token0BalInUsd float64 `json:"token0_bal_in_usd"`
	Token1BalInUsd float64 `json:"token1_bal_in_usd"`
}
