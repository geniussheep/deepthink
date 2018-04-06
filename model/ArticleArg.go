package model

type ArticleArg struct {
	PageArg
	Chid int64 `form:"chid" json:"chid"`
}
