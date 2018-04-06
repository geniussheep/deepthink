package controller

import (
	"github.com/winlion/deepthink/model"
	"github.com/winlion/deepthink/restgo"

	"github.com/gin-gonic/gin"

	"github.com/winlion/deepthink/entity"
	"github.com/winlion/deepthink/service"

	"github.com/gin-gonic/gin/binding"
)

type ArticleController struct {
	restgo.Controller
}

var articleService service.ArticleService

//初始化权限资源
func (ctrl *ArticleController) init() {
	auth := resourceService.All()
	tmp := make(map[string]int64)
	for _, a := range auth {
		tmp[a.Patern] = a.ID
	}
	restgo.AllAuth(tmp)
}

//
func (ctrl *ArticleController) Router(router *gin.Engine) {
	ctrl.init()
	r := router.Group("article")
	r.POST("create", ctrl.create)
	r.POST("update", ctrl.update)
	r.POST("search", ctrl.search)
}

//基于全部的搜索
func (ctrl *ArticleController) create(ctx *gin.Context) {
	var res entity.Article
	ctx.ShouldBindWith(&res, binding.FormPost)

	ret, err := articleService.Create(res)
	if err != nil {
		restgo.ResultFail(ctx, err.Error())
	} else {
		ctrl.init()
		restgo.ResultOkMsg(ctx, ret, "文章发布成功")
	}
}

//基于全部的搜索
func (ctrl *ArticleController) update(ctx *gin.Context) {
	var res entity.Article
	ctx.ShouldBindWith(&res, binding.FormPost)
	ret, err := articleService.Update(res)
	if err != nil {
		restgo.ResultFail(ctx, err.Error())
	} else {
		ctrl.init()
		restgo.ResultOkMsg(ctx, ret, "资源添加成功")
	}
}

//基于全部的搜索
func (ctrl *ArticleController) search(ctx *gin.Context) {
	var arg model.ArticleArg
	ctx.ShouldBindWith(&arg, binding.FormPost)
	ret := articleService.Query(arg)
	//最后响应数据列表到前端
	restgo.ResultList(ctx, ret, int64(len(ret)))
}
