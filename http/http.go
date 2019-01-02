package http

import (
	"github.com/astaxie/beego"

	"github.com/shanghai-edu/ldaproam/g"
)

func Start() {
	beego.BConfig.Listen.EnableHTTPS = true
	beego.BConfig.Listen.HTTPSAddr = g.Config().Http.Listen
	beego.BConfig.Listen.HTTPSPort = g.Config().Http.Port
	beego.BConfig.Listen.HTTPSCertFile = g.Config().Credentials.Cert
	beego.BConfig.Listen.HTTPSKeyFile = g.Config().Credentials.Key

	beego.BConfig.CopyRequestBody = true

	if !g.Config().Http.Debug {
		beego.SetLevel(beego.LevelInformational)
	}
	ConfigRouters()
	beego.Run()
}
