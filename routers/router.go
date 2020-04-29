package routers

import (
	"beego_blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{}, "*:Index")
    // /index2.html
    beego.Router("/index:page:int.html", &controllers.MainController{}, "*:Index")
}
