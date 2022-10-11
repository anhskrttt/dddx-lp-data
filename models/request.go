package models

type LPRequest struct {
	UserAddress  string `form:"user_address"`
	ProtocolId   string `form:"protocol_id"` // Default: DDDX: "dddx"
	GaugeAddress string `form:"gauge_address"`
}
