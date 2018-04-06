package service

import (
	"errors"

	"github.com/winlion/deepthink/entity"
	"github.com/winlion/deepthink/model"
	"github.com/winlion/deepthink/restgo"
)

type ArticleService struct{}

//
func (service *ArticleService) Query(arg model.ArticleArg) []entity.Article {
	var objs []entity.Article = make([]entity.Article, 0)
	orm := restgo.OrmEngin()
	t := orm.Where("1=1")
	if 0 < len(arg.Kword) {
		kword := "%" + arg.Kword + "%"
		t = t.Where("title like ? or memo like ? or tag like ?", kword, kword, kword)
	}
	if 0 != arg.Chid {

		t = t.Where("chid = ?", arg.Chid)
	}
	t.Limit(arg.GetPageFrom()).Find(&objs)
	return objs
}

//
func (service *ArticleService) Update(arg entity.Article) (int64, error) {
	orm := restgo.OrmEngin()
	return orm.ID(arg.ID).Update(&arg)

}

//
func (service *ArticleService) Create(article entity.Article) (ar entity.Article, err error) {

	if 0 == len(article.Title) {
		err = errors.New("请输入标题")

		return
	}
	if 0 == len(article.Memo) {
		err = errors.New("请输入描述")

		return
	}

	if 0 == len(article.Content) {
		err = errors.New("请输入文章内容")
		return
	}
	orm := restgo.OrmEngin()
	id, err := orm.InsertOne(&article)
	ar.ID = id
	return ar, nil
}

//
func (service *ArticleService) Deleted(arg entity.Article) {
	orm := restgo.OrmEngin()
	orm.ID(arg.ID).Update(&arg, "deleted", 1)
}
