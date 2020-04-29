package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")

	//"root:123456@tcp(127.0.0.1:3306)/beego_blog?charset=utf8"
	dburl := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"

	//set default database
	orm.RegisterDataBase("default", "mysql",dburl ,30)

	//register model
	orm.RegisterModel(new(Link), new(Mood), new(Post), new(Tag), new(TagPost), new(User))
}

//查询最新的4篇文章
func GetLatestBlog() []*Post {
	post:= Post{}
	//从文章表中过滤出状态正常的文章
	query := orm.NewOrm().QueryTable(&post).Filter("status", 0)
	//查询数量
	count, _ := query.Count()
	var result []*Post
	if count >0{
		query.OrderBy("-posttime").Limit(4).All(&result)
	}
	return result
}

//查询最新的4篇文章
func GetHotBlog() []*Post {
	post:= Post{}
	//从文章表中过滤出状态正常的文章
	query := orm.NewOrm().QueryTable(&post).Filter("status", 0)
	//查询数量
	count, _ := query.Count()
	var result []*Post
	if count >0{
		query.OrderBy("-views").Limit(4).All(&result)
	}
	return result
}

//友情链接
func GetLinks() []*Link {
	link := Link{}
	query := orm.NewOrm().QueryTable(&link)
	count, _ := query.Count()
	var result []*Link
	if count >0 {
		query.OrderBy("-rank").All(&result)
	}
	return result
}