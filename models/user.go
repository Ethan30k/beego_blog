package models

type User struct {
	Id int
	Username string `orm:"size(15)"`
	Password string `orm:"size(50)"`
	Email string `orm:"size(50)"`
	Logincount int //登录次数
	Authkey string `orm:"size(10)"`
	Active int //是否激活
}

