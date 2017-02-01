package controllers

import (
	"github.com/astaxie/beego"
	"blog/models"
)

type HomeController struct {
	beego.Controller
}

func (p *HomeController) Get(){
	p.Data["IsHome"] = true
	p.Data["Categories"] = models.GetAllCategories()
	p.Data["Topics"] = models.GetAllTopics()
	p.TplName = "home.html"
}