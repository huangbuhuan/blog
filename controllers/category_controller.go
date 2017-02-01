package controllers

import (
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (p *CategoryController) Get() {
	p.Data["IsTopic"] = true
}