package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (p *LoginController) Get() {
	p.TplName = "login.html"
}

func (p *LoginController) Post() {

	if beego.AppConfig.String("admin") == p.Input().Get("username") &&
		beego.AppConfig.String("passwd") == p.Input().Get("passwd") {
		max := 0
		if p.Input().Get("autoLogin") == "on" {
			max = 3600
		}
		p.Ctx.SetCookie("name", p.Input().Get("username"), max, "/")
		fmt.Println("login")
	} else {
		fmt.Println("error")
	}
	p.Redirect("/", 301)
	return
}