package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type TagPost struct {
	Id int
	Tagid int //标签id
	Postid int //文章id
	Poststatus int //文章状态
	Posttime time.Time //文章发表时间
}

func (tagpost *TagPost)TableName() string {
	//从配置文件中获取表的前缀
	dbprefix := beego.AppConfig.String("dbprefix")
	return dbprefix + "tagpost"
}

//插入
func (tagpost *TagPost) Insert() error{
	if _, err := orm.NewOrm().Insert(tagpost);err != nil{
		return err
	}
	return nil
}

//读取
func (tagpost *TagPost)Read(fields ...string) error {
	if err := orm.NewOrm().Read(tagpost, fields...);err !=nil{
		return err
	}
	return nil
}

//删除
func (tagpost *TagPost)Delete() error {
	if _, err := orm.NewOrm().Delete(tagpost);err !=nil{
		return err
	}
	return nil
}

//更新
func (tagpost *TagPost)Update(fields ...string) error {
	if _,err := orm.NewOrm().Update(tagpost, fields...);err !=nil{
		return err
	}
	return nil
}