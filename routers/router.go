package routers

import (
	"beego_blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{}, "*:Index")
    // /index2.html
    beego.Router("/index:page:int.html", &controllers.MainController{}, "*:Index")

    //文章详情
    //  /article/2
    beego.Router("/article/:id:int", &controllers.MainController{},"*:Show")

    //成长录
    beego.Router("/life.html", &controllers.MainController{},"*:Life")
    //成长录分页
    beego.Router("/life:page:int.html", &controllers.MainController{},"*:Life")

    //关于我
    beego.Router("/about.html", &controllers.MainController{},"*:About")
}
