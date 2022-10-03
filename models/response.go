package models

type LPInfo struct {
	UserAddress string           `json:"user_address"`
	LPBalance   TokenPairBalance `json:"lp_balance"`
}

type GetLpOfProtocolResponse struct {
	UserAddress string           `json:"user_address"`
	Token0 TokenBalance `json:"token0"`
	Token1 TokenBalance `json:"token1"`
	ProtocolId  string           `json:"protocol_id"`
	Pool        PoolSimple       `json:"pool"`
}
