package controller

import (
	"github.com/winlion/deepthink/model"
	"github.com/winlion/deepthink/restgo"

	"github.com/gin-gonic/gin"

	"github.com/winlion/deepthink/entity"
	"github.com/winlion/deepthink/service"

	"github.com/gin-gonic/gin/binding"
)

type ChannelController struct {
	restgo.Controller
}

var channelService service.ChannelService

//
func (ctrl *ChannelController) Router(router *gin.Engine) {

	r := router.Group("channel")
	r.POST("create", ctrl.create)
	r.POST("update", ctrl.update)
	r.POST("search", ctrl.search)
}

//基于全部的搜索
func (ctrl *ChannelController) create(ctx *gin.Context) {
	var res entity.Channel
	ctx.ShouldBindWith(&res, binding.FormPost)

	ret, err := channelService.Create(res)
	if err != nil {
		restgo.ResultFail(ctx, err.Error())
	} else {

		restgo.ResultOkMsg(ctx, ret, "文章发布成功")
	}
}

//基于全部的搜索
func (ctrl *ChannelController) update(ctx *gin.Context) {
	var res entity.Channel
	ctx.ShouldBindWith(&res, binding.FormPost)
	ret, err := channelService.Update(res)
	if err != nil {
		restgo.ResultFail(ctx, err.Error())
	} else {

		restgo.ResultOkMsg(ctx, ret, "频道添加成功")
	}
}

//基于全部的搜索
func (ctrl *ChannelController) search(ctx *gin.Context) {
	var arg model.ChannelArg
	ctx.ShouldBindWith(&arg, binding.FormPost)
	ret := channelService.Query(arg)
	//最后响应数据列表到前端
	restgo.ResultList(ctx, ret, int64(len(ret)))
}
