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

func (this *MainController)Prepare()  {
	var page int
	var err error
	page, err = strconv.Atoi(this.Ctx.Input.Param(":page"))
	if err!=nil{
		page =1
		fmt.Println(err)
	}
	this.Pager = models.NewPager(page,3,0,"")
}

//首页
func (this *MainController) Index() {
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
	this.setRight()
	this.display("index")
	this.setHeadMeater()
}

//设置首页右侧部分
func (this *MainController) setRight() {
	//最新文章
	this.Data["latestblog"] = models.GetLatestBlog()
	//浏览量最多的4篇文章
	this.Data["hotblog"] = models.GetHotBlog()
	//友情链接
	this.Data["links"] = models.GetLinks()
}

func (this *MainController) display(tplname string) {
	this.Layout = "double/layout.html"
	this.TplName = "double/"+tplname+".html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["head"] = "double/head.html"
	this.LayoutSections["foot"] = "double/foot.html"

	if tplname == "index"{
		this.LayoutSections["banner"] = "double/banner.html"
		this.LayoutSections["middle"] = "double/middle.html"
		this.LayoutSections["right"] = "double/right.html"
	}else if tplname == "life"{
		this.LayoutSections["right"] = "double/right.html"
	}
}

//设置头部信息
func (this *MainController)setHeadMeater()  {
	this.Data["title"] = beego.AppConfig.String("title")
	this.Data["keywords"] = beego.AppConfig.String("keywords")
	this.Data["description"] = beego.AppConfig.String("description")
}

//通过文章id查看文章详情
func (this *MainController)Show()  {
	//获取文章id并转换整数
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err!= nil{
		this.Redirect("/404",302)
	}
	//创建文章结构体
	post := new(models.Post)
	post.Id = id
	//查询文章
	err = post.Read()
	if err != nil{
		this.Redirect("/404",302)
	}
	//浏览量+1
	post.Views++
	//更新浏览量
	post.Update("Views")

	this.Data["post"] = post
	//获取上一篇文章与下一篇文章
	pre, next := post.GetPreAndNext()
	this.Data["pre"] = pre
	this.Data["next"] = next
	this.display("article")
	this.Data["smalltitle"] = "文章详情"
}

func (this *MainController)Life() {
	var list []*models.Post
	post := models.Post{}
	//获得文章表的句柄，并设置过滤条件（正常状态的文章）
	query := orm.NewOrm().QueryTable(&post).Filter("status", 0)

	//获得符合条件的记录数
	count, _ := query.Count()

	//设置总的数量
	this.Pager.SetTotalnum(int(count))
	//设置每页对应的路径
	this.Pager.SetUrlpath("/life%d.html")

	if count >0 {
		offset := (this.Pager.Page - 1) * this.Pager.Pagesize
		_, err := query.OrderBy("-istop", "-views").Limit(this.Pager.Pagesize, offset).All(&list)
		if err != nil{
			fmt.Println("err=",err)
		}
	}

	this.Data["list"] = list
	this.Data["pagebar"] = this.Pager.ToString()
	this.setRight()
	this.display("life")
	this.setHeadMeater()
}

func (this *MainController) About() {
	this.setHeadMeater()
	this.display("about")
}