package admin

import (
	"beego_blog/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type baseController struct {
	beego.Controller
	userid int
	username       string
	controllerName string
	actionName     string
	pager *models.Pager
}

func (this *baseController) Prepare() {
	//获取控制器名称和方法名称
	this.controllerName, this.actionName = this.GetControllerAndAction()
	//去除控制器名称尾部的Controller并将结果转化为小写
	this.controllerName = strings.ToLower(this.controllerName[:len(this.controllerName)-10])
	//将方法名称转化为小写
	this.actionName = strings.ToLower(this.actionName)

	this.auth()
	page, err := this.GetInt("page")
	if err != nil{
		page =1
	}
	pagesize :=2
	this.pager = models.NewPager(page,pagesize,0,"")
}

func (this *baseController) auth() {
	if this.controllerName == "main" || (this.controllerName == "account" && this.actionName == "login") {
		return
	}

	arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
	if len(arr) == 2 {
		idstr, password :=arr[0], arr[1]
		//将id转换为整数
		id, _ := strconv.Atoi(idstr)
		if id >0{
			user := new(models.User)
			user.Id =id
			if user.Read() == nil &&user.Password == password{
				this.userid = user.Id
				this.username = user.Username
			}
		}
	}
	if this.userid == 0{
		this.Redirect("/admin/login", 302)
	}
}

func (this *baseController) display(tplname ...string) {
	modileName := "admin/"
	this.Layout = modileName + "layout.html"

	this.Data["version"] = beego.AppConfig.String("version")
	this.Data["adminname"] = this.username

	if len(tplname) == 1 {
		this.TplName = modileName + tplname[0] + ".html"
	} else {
		this.TplName = modileName + this.controllerName + "_" + this.actionName + ".html"
	}
}

func (this *baseController)showmsg(msg ...string)  {
	if len(msg) ==0{
		msg = append(msg, "出错了")
	}
	//拼接上一个页面的地址
	msg = append(msg, this.Ctx.Request.Referer())
	this.Data["msg"] = msg[0]
	this.Data["redirect"] = msg[1]
	this.display("showmsg")
	this.Render()
	this.StopRun()
}