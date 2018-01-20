package main

import (
	_ "github.com/gaodongxu1/Blog/application/blog/routers/admin"
	_ "github.com/gaodongxu1/Blog/application/blog/routers/api"
	"github.com/astaxie/beego"
	mysql "github.com/gaodongxu1/Blog/application/blog/orm"
    "github.com/gaodongxu1/Blog/application/blog/filters"

	"github.com/astaxie/beego/orm"
	"github.com/gaodongxu1/Blog/application/blog/models/AutomaticCreated"	
)

func init() {
 orm.RegisterModel(new(models.Student))
}

func createTable() {
 name := "Blog"
 force := true
 verbose := true
 err := orm.RunSyncdb(name, force, verbose)
 if err != nil {
  beego.Error(err)
 }
}

func main() {
 mysql.InitMysql()
 o := orm.NewOrm()
 o.Using("Blog")
 createTable()
 beego.InsertFilter("/*",beego.BeforeRouter,filters.LoginFilter)
 beego.Run()
}