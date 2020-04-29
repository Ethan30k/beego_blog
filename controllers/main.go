package controllers

import (
	"beego_blog/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type MainController struct {
	beego.Controller
	Pager *models.Pager
}

//首页
func (this *MainController) Index() {
	var page int
	var err error
	page, err = strconv.Atoi(this.Ctx.Input.Param(":page"))
	fmt.Println("this.Ctx.Input.Param:",page)
	if err!=nil{
		page =1
		fmt.Println(err)
	}


	this.Pager = models.NewPager(page,2,0,"")

	var list []*models.Post
	post := models.Post{}
	//获得文章表的句柄，并设置过滤条件（正常状态的文章）
	query := orm.NewOrm().QueryTable(&post).Filter("status", 0)

	//获得符合条件的记录数
	count, _ := query.Count()

	//设置总的数量
	this.Pager.SetTotalnum(int(count))
	//设置每页对应的路径
	this.Pager.SetUrlpath("/index%d.html")

	if count >0 {
		offset := (this.Pager.Page - 1) * this.Pager.Pagesize
		_, err := query.OrderBy("-istop", "-views").Limit(this.Pager.Pagesize, offset).All(&list)
		if err != nil{
			fmt.Println("err=",err)
		}
	}

	this.Data["list"] = list
	this.Data["pagebar"] = this.Pager.ToString()
	//最新文章
	this.Data["latestblog"] = models.GetLatestBlog()
	//浏览量最多的4篇文章
	this.Data["hotblog"] = models.GetHotBlog()
	//友情链接
	this.Data["links"] = models.GetLinks()

	this.Layout = "double/layout.html"
	this.TplName = "double/index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["head"] = "double/head.html"
	this.LayoutSections["banner"] = "double/banner.html"
	this.LayoutSections["middle"] = "double/middle.html"
	this.LayoutSections["right"] = "double/right.html"
	this.LayoutSections["foot"] = "double/foot.html"
}
