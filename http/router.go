package http

import (
	"github.com/astaxie/beego"
	"github.com/shanghai-edu/ldaproam/http/controllers"
)

func ConfigRouters() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1/bind", &controllers.BindController{})
	beego.Router("/api/v1/search", &controllers.SearchController{})
}
