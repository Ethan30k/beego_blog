package models

type Tag struct {
	Id int
	Name string `orm:"size(20)"` //标签名称
	Count int //文章数量
}