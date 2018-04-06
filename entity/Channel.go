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
type Channel struct {
	ID       int64           `xorm:"pk autoincr 'id'" form:"id" json:"id"`
	Pid      int64           `xorm:"bigint" form:"pid" json:"pid"`
	Name     string          `xorm:"varchar(140)" form:"title" json:"title"`
	Memo     string          `xorm:"varchar(180)" form:"memo" json:"memo"`
	CreateAt restgo.JsonTime `xorm:"created" form:"createat" json:"createat"  time_format:"2006-01-02 15:04:05"`
	Deleted  int             `xorm:"int" form:"deleted" json:"deleted"`
}
