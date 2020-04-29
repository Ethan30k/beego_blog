package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

//文章
type Post struct {
	Id int
	Userid int
	//作者
	Author string `orm:"size(15)"`
	//标题
	Title string `orm:"size(100)"`
	//标题颜色
	Color string `orm:"size(7)"`
	//文章内容
	Content string `orm:"type(text)"`
	//标签名称
	Tags string `orm:"size(100)"`
	//浏览量
	Views int
	//状态
	Status int
	//发表时间
	Posttime time.Time `orm:"type(datetime)"`
	//是否置顶
	Istop int
	//封面
	Cover string `orm:"size(70)"`
}

func (post *Post)TableName() string {
	//从配置文件中获取表的前缀
	dbprefix := beego.AppConfig.String("dbprefix")
	return dbprefix + "post"
}

//插入
func (post *Post) Insert() error{
	if _, err := orm.NewOrm().Insert(post);err != nil{
		return err
	}
	return nil
}

//读取
func (post *Post)Read(fields ...string) error {
	if err := orm.NewOrm().Read(post, fields...);err !=nil{
		return err
	}
	return nil
}

//更新
func (post *Post)Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(post, fields...);err !=nil{
		return err
	}
	return nil
}

func (post *Post)TagsLink() string {
	if post.Tags == "" {
		return ""
	}
	tagslink := strings.Trim(post.Tags, ",")
	return tagslink
}

func (post *Post)Link() string {
	return "/article/"+strconv.Itoa(post.Id)
}

func (post *Post)ColorTitle() string {
	if post.Color != ""{
		return fmt.Sprintf("<span style='color:%s'>%s</span>", post.Color, post.Title)
	}
	return post.Title
}

func (post *Post)Excerpt() string {
	return post.Content
}

//根据当前文章获取上一篇与下一篇文章
func (this *Post)GetPreAndNext() (pre, next *Post) {
	//上一篇文章
	pre = &Post{}
	err := orm.NewOrm().QueryTable(new(Post)).OrderBy("-id").Filter("id__lt", this.Id).Filter("status", 0).Limit(1).One(pre)
	if err!=nil{
		pre=nil
	}

	//下一篇文章
	next = &Post{}
	err = orm.NewOrm().QueryTable(new(Post)).OrderBy("id").Filter("id__gt", this.Id).Filter("status", 0).Limit(1).One(next)
	if err!=nil{
		next=nil
	}
	return
}