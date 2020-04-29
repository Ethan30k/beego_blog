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
}
