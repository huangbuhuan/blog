package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
)

func init()  {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("dbconf"))
	fmt.Printf("数据库连接成功！\n", )
}

func main() {
	orm.Debug = true
	beego.Run()
}

