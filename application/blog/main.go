package main

import (
	_ "github.com/gaodongxu1/Blog/application/blog/routers/api"
	"github.com/astaxie/beego"
	"github.com/gaodongxu1/Blog/application/blog/orm"
)

func main() {
	orm.InitMysql() //调用注册的数据库(orm)
	beego.Run()
}

