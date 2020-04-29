package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Tag struct {
	Id int
	Name string `orm:"size(20)"` //标签名称
	Count int //文章数量
}

func (tag *Tag)TableName() string {
	//从配置文件中获取表的前缀
	dbprefix := beego.AppConfig.String("dbprefix")
	return dbprefix + "tag"
}

//插入
func (tag *Tag) Insert() error{
	if _, err := orm.NewOrm().Insert(tag);err != nil{
		return err
	}
	return nil
}

//读取
func (tag *Tag)Read(fields ...string) error {
	if err := orm.NewOrm().Read(tag, fields...);err !=nil{
		return err
	}
	return nil
}

//更新
func (tag *Tag)Update(fields ...string) error {
	if err := orm.NewOrm().Read(tag, fields...);err !=nil{
		return err
	}
	return nil
}