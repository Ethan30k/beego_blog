package admin

import (
	"beego_blog/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"time"
)

type MoodController struct {
	baseController
}

//碎言碎语
func (this *MoodController) List() {
	var list []*models.Mood
	query := orm.NewOrm().QueryTable(new(models.Mood))
	count, _ := query.Count()
	if count > 0 {
		offset := (this.pager.Page - 1) * this.pager.Pagesize
		query.OrderBy("-id").Limit(this.pager.Pagesize, offset).All(&list)
	}
	this.Data["list"] = list
	this.pager.SetTotalnum(int(count))
	this.pager.SetUrlpath("/admin/mood/list?page=%d")
	this.Data["pagebar"] = this.pager.ToString()

	this.display()
}

//删除闲言碎语
func (this *MoodController) Delete() {
	//获取id
	id, err := this.GetInt("id")
	if err != nil {
		this.showmsg("删除失败")
	}
	mood := models.Mood{Id: id}
	if err = mood.Read(); err == nil {
		mood.Delete()
	}
	this.Redirect("/admin/mood/list", 302)
}

//添加闲言碎语
func (this *MoodController) Add() {
	//判断是否是POST请求
	if this.Ctx.Request.Method == "POST"{
		//获取碎言碎语的内容
		content := this.GetString("content")
		var mood models.Mood
		mood.Content = content
		//初始化随机数种子
		rand.Seed(time.Now().Unix())
		var r = rand.Intn(11)
		mood.Cover="/static/upload/blog" + fmt.Sprintf("%d",r)+".jpg"
		mood.Posttime = time.Now()
		if err:=mood.Insert();err !=nil{
			this.showmsg(err.Error())
		}

		this.Redirect("/admin/mood/list", 302)
	}
	this.display()

}
