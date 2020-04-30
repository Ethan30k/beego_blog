package admin

import (
	"beego_blog/models"
	"strconv"
	"strings"
)

type AccountController struct {
	baseController
}

func (this *AccountController)Login()  {
	//判断请求方式是否是post
	if this.GetString("dosubmit") == "yes"{
		//获取账号并去除两边空格
		account := strings.TrimSpace(this.GetString("account"))
		password := strings.TrimSpace(this.GetString("password"))
		remember := strings.TrimSpace(this.GetString("remember"))
		if account !="" &&password !=""{
			var user = &models.User{}
			user.Username = account
			//根据账号查询用户，并且判断查询到的密码和用户输入的密码通过MD5哈希之后的结果是否一致
			if user.Read("username") != nil||user.Password!=models.MD5([]byte(password)){
				this.Data["errmsg"] = "账号或密码错误！"
			}else if user.Active==0{//判断该账号是否激活
				this.Data["errmsg"] = "账号未激活！"
			}else {
				//登录次数+1
				user.Logincount+=1
				//更新Logincount
				user.Update("logincount")
				authkey := models.MD5([]byte(password))
				if remember == "yes"{
					this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 60*69*24*7)
				}else{
					this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey)
				}
				this.Redirect("/admin", 302)
			}
		}
	}
	//重定向到后台首页
	this.TplName="admin/" + this.controllerName + "_" + this.actionName + ".html"
}
