package model

type ChannelArg struct {
	PageArg
	Deleted int `form:"deleted" json:"deleted"`
}
