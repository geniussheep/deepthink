package service

import (
	"errors"

	"github.com/winlion/deepthink/entity"
	"github.com/winlion/deepthink/model"
	"github.com/winlion/deepthink/restgo"
)

type ChannelService struct{}

//
func (service *ChannelService) Query(arg model.ChannelArg) []entity.Channel {
	var objs []entity.Channel = make([]entity.Channel, 0)
	orm := restgo.OrmEngin()
	t := orm.Where("1=1")
	if 0 < len(arg.Kword) {
		t = t.Where("name like ? ", "%"+arg.Kword+"%")
	}

	t = t.Where(" deleted = ?", arg.Deleted)

	t.Limit(arg.GetPageFrom()).Find(&objs)
	return objs
}

//
func (service *ChannelService) Update(arg entity.Channel) (int64, error) {
	orm := restgo.OrmEngin()
	return orm.ID(arg.ID).Update(&arg, "name", arg.Name)
}

//
func (service *ChannelService) Create(arg entity.Channel) (ch entity.Channel, err error) {
	ch = arg
	if 0 == len(arg.Name) {
		err = errors.New("缺少频道名称")
		return
	}
	orm := restgo.OrmEngin()
	id, err := orm.InsertOne(&arg)
	ch.ID = id

	return ch, err
}

//
func (service *ChannelService) Deleted(arg entity.Channel) {
	orm := restgo.OrmEngin()
	orm.ID(arg.ID).Update(&arg, "deleted", 1)
}
