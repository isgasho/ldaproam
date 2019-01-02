package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shanghai-edu/ldaproam/g"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.Output.Body([]byte("ldaproam, version " + g.VERSION))
}
