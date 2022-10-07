package models

import "math/big"

// LP & Farming response
type AllLPBalResponse struct {
	Token0      TokenBalance `json:"token0"`
	Token1      TokenBalance `json:"token1"`
	Pool        PoolSimple   `json:"pool"`
}

// LP & Farming response
type LPFarmResponse struct {
	UserAddress string       `json:"user_address"`
	Token0      TokenBalance `json:"token0"`
	Token1      TokenBalance `json:"token1"`
	ProtocolId  string       `json:"protocol_id"`
	Pool        PoolSimple   `json:"pool"`
}

type UserInformationLocked struct {
	Amount *big.Int
	End    *big.Int
}

type StakeResponse struct {
	UserAddress    string                  `json:"user_address"`
	TokenSymbol    string                  `json:"token"`
	Staked_balance []UserInformationLocked `json:"staked_balance"`
}
