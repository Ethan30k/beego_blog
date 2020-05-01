package routers

import (
	"beego_blog/controllers"
	"beego_blog/controllers/admin"
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

    beego.Router("/mood.html", &controllers.MainController{},"*:Mood")

    /*-----------------------------------后台页面-------------------------------------------*/
	//首页
    beego.Router("/admin",&admin.IndexController{}, "*:Index")

    /*-----------------------------------账户管理------------------------------------------------*/
	//登陆
	beego.Router("/admin/login", &admin.AccountController{}, "*:Login")
	beego.Router("/admin/logout", &admin.AccountController{}, "*:Logout")
	beego.Router("/admin/account/profile", &admin.AccountController{}, "*:Profile")

	/*-----------------------------------说说管理------------------------------------------------*/
	beego.Router("/admin/mood/list", &admin.MoodController{}, "*:List")
	beego.Router("/admin/mood/delete", &admin.MoodController{}, "*:Delete")
	beego.Router("/admin/mood/add", &admin.MoodController{}, "*:Add")

	/*-----------------------------------用户管理--------------------------------------------------*/
	beego.Router("/admin/user/list", &admin.UserController{}, "*:List")
	beego.Router("/admin/user/delete", &admin.UserController{}, "*:Delete")
	beego.Router("/admin/user/edit", &admin.UserController{}, "*:Edit")
	beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")

	/*------------------------------------友链管理--------------------------------------------------*/
	beego.Router("/admin/link/list", &admin.LinkController{}, "*:List")
	beego.Router("/admin/link/edit", &admin.LinkController{}, "*:Edit")
	beego.Router("/admin/link/delete", &admin.LinkController{}, "*:Delete")
	beego.Router("/admin/link/add", &admin.LinkController{}, "*:Add")

	/*-----------------------------------文章管理---------------------------------------------------*/
	beego.Router("/admin/article/list", &admin.ArticleController{}, "*:List")
	beego.Router("/admin/article/add", &admin.ArticleController{}, "*:Add")
	beego.Router("/admin/article/save", &admin.ArticleController{}, "*:Save")
	beego.Router("/admin/article/delete", &admin.ArticleController{}, "*:Delete")
	beego.Router("/admin/article/batch", &admin.ArticleController{}, "*:Batch")
}
