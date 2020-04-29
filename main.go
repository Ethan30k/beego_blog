package main

import (
	_ "beego_blog/models"
	_ "beego_blog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
