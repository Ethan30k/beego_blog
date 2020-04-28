package models

//友情链接
type Link struct {
	Id int
	Sitename string `orm:"size(80)"` //网站名称
	Url string `orm:"size(200)"` //网站
	Rank int //排序
} 
