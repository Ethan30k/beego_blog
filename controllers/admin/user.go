package admin

import (
	"beego_blog/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strings"
)

type UserController struct {
	baseController
}

func (this *UserController) List() {
	var list []models.User
	query := orm.NewOrm().QueryTable(new(models.User))
	count, _ := query.Count()
	if count >0 {
		offset:=(this.pager.Page-1)*this.pager.Pagesize
		query.OrderBy("-id").Limit(this.pager.Pagesize, offset).All(&list)
	}
	this.Data["list"] = list
	this.pager.SetTotalnum(int(count))
	this.pager.SetUrlpath("/admin/user/list?page=%d")
	this.Data["pagebar"] = this.pager.ToString()
	this.display()
}

func (this *UserController) Delete() {
	//获取id
	id, err := this.GetInt("id")
	if id ==7 {
		this.showmsg("不能删除超级管理员")
	}
	if err != nil {
		this.showmsg("删除失败")
	}
	user := &models.User{Id: id}
	if err = user.Read(); err == nil {
		user.Delete()
	}
	this.Redirect("/admin/user/list", 302)
}

func (this *UserController) Edit() {
	//获取用户id
	id, _ := this.GetInt("id")
	user := &models.User{Id: id}
	if err:=user.Read();err!=nil{
		this.showmsg("用户不存在")
	}
	fmt.Println(user)
	//用户存储错误提示
	errmsg := make(map[string]string)
	if this.Ctx.Request.Method=="POST"{
		password := strings.TrimSpace(this.GetString("password"))
		password2 := strings.TrimSpace(this.GetString("password2"))
		email := strings.TrimSpace(this.GetString("email"))
		active,_ := this.GetInt("active")
		valid := validation.Validation{}

		if password !=""{
			if result:=valid.Required(password2, "password2");!result.Ok{
				errmsg["password2"] = "确认密码不能为空！"
			}else if password!=password2{
				errmsg["password2"] = "两次输入的密码不一致"
			}else {
				user.Password = models.MD5([]byte(password))
			}
		}
		if result := valid.Required(email, "email");!result.Ok{
			errmsg["email"] = "邮箱不能为空！"
		}else if result := valid.Email(email, "email");!result.Ok{
			errmsg["email"] = "邮箱不合法！"
		}else {
			user.Email=email
		}

		if active>0{
			user.Active = 1
		}else {
			user.Active = 0
		}

		if len(errmsg)==0{
			user.Update()
			this.Redirect("/admin/user/list", 302)
		}
	}
	this.Data["user"] = user
	this.Data["errmsg"] = errmsg
	this.display()
}