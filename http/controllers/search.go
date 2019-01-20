package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/shanghai-edu/ldaproam/backend"
	"github.com/shanghai-edu/ldaproam/cron"
	"github.com/shanghai-edu/ldaproam/g"
	"github.com/shanghai-edu/ldaproam/gocrypto"
	"github.com/shanghai-edu/ldaproam/metadata"
)

type SearchController struct {
	beego.Controller
}

type SearchResult struct {
	Success bool                  `json:"success"`
	Result  []backend.LDAP_RESULT `json:"result"`
}

func (this *SearchController) Post() {
	var req g.HttpSearchReq
	var msgResult MsgResult
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

	res, err := backend.LDAP_SearchUser(req.Body.Data.Username, backend.TranslateAttributes(req.Body.Data.Attributes, g.Config().Backend.AttributesMap))
	if err != nil {
		msgResult.Msg = err.Error()
		this.Data["json"] = msgResult
		this.ServeJSON()
		return
	}
	for i, r := range res {
		if _, ok := cron.TrustDN.Load(r.DN); ok {
			continue
		} else {
			EmptyMap := make(map[string][]string)
			res[i].Attributes = EmptyMap
		}
	}
	var searchResult SearchResult
	searchResult.Result = res
	searchResult.Success = true
	this.Data["json"] = searchResult
	this.ServeJSON()
}
