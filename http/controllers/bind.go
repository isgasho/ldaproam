package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/shanghai-edu/ldaproam/backend"
	"github.com/shanghai-edu/ldaproam/cron"
	"github.com/shanghai-edu/ldaproam/g"
	"github.com/shanghai-edu/ldaproam/gocrypto"
	"github.com/shanghai-edu/ldaproam/metadata"
)

type BindController struct {
	beego.Controller
}

type MsgResult struct {
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}

func (this *BindController) Post() {
	var req g.HttpBindReq
	var msgResult MsgResult
	g.DebugLog(string(this.Ctx.Input.RequestBody))
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		msgResult.Msg = err.Error()
		this.Data["json"] = msgResult
		this.ServeJSON()
		return
	}
	domainName := req.Body.From
	m, err := metadata.GetMetadataByDomainName(domainName)
	if err != nil {
		msgResult.Msg = err.Error()
		this.Data["json"] = msgResult
		this.ServeJSON()
		return
	}
	publicKey, err := gocrypto.GetPublicFromCert([]byte(m.Certificate))
	if err != nil {
		msgResult.Msg = err.Error()
		this.Data["json"] = msgResult
		this.ServeJSON()
		return
	}
	js, _ := json.Marshal(req.Body)
	err = gocrypto.RsaVerify(string(js), req.Sign, publicKey)
	if err != nil {
		msgResult.Msg = err.Error()
		this.Data["json"] = msgResult
		this.ServeJSON()
		return
	}
	password := req.Body.Data.Password
	decryptStr, err := gocrypto.RsaDecrypt(password, g.Config().Credentials.PrivateKey)
	if err == nil {
		password = decryptStr
	}

	err = backend.LDAP_Bind(req.Body.Data.Dn, password)
	if err != nil {
		msgResult.Msg = err.Error()
		this.Data["json"] = msgResult
		this.ServeJSON()
		return
	}
	err = cron.AddTrustDN(req.Body.Data.Dn)
	if err != nil {
		msgResult.Msg = err.Error()
		this.Data["json"] = msgResult
		this.ServeJSON()
		return
	}
	msgResult.Msg = "Ldap Bind Success"
	msgResult.Success = true
	this.Data["json"] = msgResult
	this.ServeJSON()
}
