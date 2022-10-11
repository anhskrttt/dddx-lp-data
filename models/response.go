package models

import "math/big"

type AllPoolsOfProtocolResponse struct {
	ProtocolId string `json:"protocol_id"`
	Length int `json:"length"`
	Pools []string `json:"pools"`

}

type AllBalOfUserResponse struct {
	UserAddress string `json:"user_address"`
	ProtocolId string  `json:"protocol_id"`
	LPBalance []AllLPFarmBalResponse `json:"lp_balance"`
	FarmBalance []AllLPFarmBalResponse `json:"farm_balance"`
	StakedBalance AllStakeResponse `json:"stakedBalance"`
}

// LP & Farming response
type AllLPFarmBalResponse struct {
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

type AllStakeResponse struct {
	TokenSymbol    string                  `json:"token"`
	Staked_balance []UserInformationLocked `json:"staked_balance"`
}
