package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (p *MainController) Get() {
	p.Data["Website"] = "buhuan.me"
	p.Data["Email"] = "astaxie@gmail.com"
	p.TplName = "index.tpl"
}
