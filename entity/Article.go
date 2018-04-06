package entity

import (
	"github.com/winlion/deepthink/restgo"
)

/*
chid
title
memo
content
tag
picture
*/
type Article struct {
	ID       int64           `xorm:"pk autoincr 'id'" form:"id" json:"id"`
	Chid     int64           `xorm:"bigint" form:"chid" json:"account"`
	Title    string          `xorm:"varchar(140)" form:"title" json:"title"`
	Memo     string          `xorm:"varchar(180)" form:"memo" json:"memo"`
	CreateAt restgo.JsonTime `xorm:"created" form:"createat" json:"createat"  time_format:"2006-01-02 15:04:05"`
	Content  string          `xorm:"text" form:"content" json:"content"`
	Tag      string          `xorm:"varchar(40)" form:"tag" json:"tag"`
	Cover    string          `xorm:"varchar(150)" form:"cover" json:"cover"`
	Deleted  int             `xorm:"deleted" json:"deleted"`
}
