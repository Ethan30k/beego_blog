package models

import "time"

//文章
type Post struct {
	Id int
	Userid int
	//作者
	Author string `orm:"size(15)"`
	//标题
	Title string `orm:"size(100)"`
	//标题颜色
	Color string `orm:"size(7)"`
	//文章内容
	Content string `orm:"type(text)"`
	//标签名称
	Tags string `orm:"size(100)"`
	//浏览量
	Views int
	//状态
	Status int
	//发表时间
	Posttime time.Time `orm:"type(datetime)"`
	//是否置顶
	Istop int
	//封面
	Cover string `orm:"size(70)"`
}
