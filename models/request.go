package models

type GetLpOfProtocolRequest struct {
	UserAddress string `form:"user_address"`
	ProtocolId  string `form:"protocol_id"` // Default: DDDX: "dddx"
	PoolAddress string `form:"pool_address"`
}


type GetFarmOfProtocolRequest struct {
	UserAddress string `form:"user_address"`
	ProtocolId  string `form:"protocol_id"` // Default: DDDX: "dddx"
	GaugeAddress string `form:"gauge_address"`
}

type GetStakedOfProtocolRequest struct {
	UserAddress string `form:"user_address"`
	VeAddress string `form:"vote_escrowed_address"`
	ProtocolId  string `form:"protocol_id"` // Default: DDDX: "dddx"
}
