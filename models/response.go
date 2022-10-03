package models

import "math/big"

type LPInfo struct {
	UserAddress string           `json:"user_address"`
	LPBalance   TokenPairBalance `json:"lp_balance"`
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
