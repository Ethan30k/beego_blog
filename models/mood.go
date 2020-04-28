package models

//碎言碎语
type Mood struct {
	Id int
	//说说内容
	Content string `orm:"type(text)"`
	//封面路径
	Cover string `orm:"size(70)"`
	//发布时间
	Posttime string `orm:"size(datetime)"`
}
