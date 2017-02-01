package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
	//首页路由注册
	beego.Router("/home", &controllers.HomeController{})
	//登录页面路由注册
	beego.Router("/login", &controllers.LoginController{})
	//分类页面路由注册
	beego.Router("/category", &controllers.CategoryController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/topic", &controllers.TopicController{})
}
