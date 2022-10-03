package models

type GetLpOfProtocolRequest struct {
	UserAddress string `form:"user_address"`
	ProtocolId  string `form:"protocol_id"` // Default: DDDX:1
	PoolAddress string `form:"pool_address"`
}
