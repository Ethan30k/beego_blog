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