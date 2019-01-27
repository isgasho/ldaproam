package controllers

import (
	"encoding/json"
	"fmt"
	"strings"

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

func getUsernameAttrFromFilter(filter string) (usernameAttr string) {
	split := strings.Split(filter, "=")
	split2 := strings.Split(split[0], "(")
	usernameAttr = split2[len(split2)-1]
	return
}

func translateAttributeMap(LdapResutlAttirbutes map[string][]string, attributeMap map[string]string) map[string][]string {
	newMap := map[string]string{}
	for k, v := range attributeMap {
		newMap[v] = k
	}

	newLdapResutlAttirbutes := map[string][]string{}
	for k, v := range LdapResutlAttirbutes {
		if newKey, ok := newMap[k]; ok {
			newLdapResutlAttirbutes[newKey] = v
		} else {
			newLdapResutlAttirbutes[k] = v
		}
	}
	return newLdapResutlAttirbutes
}

func (this *SearchController) Post() {
	var req g.HttpSearchReq
	var msgResult MsgResult
	g.DebugLog(fmt.Sprintf("Request on /api/v1/search, req: %s", string(this.Ctx.Input.RequestBody)))
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
	usernameAttr := getUsernameAttrFromFilter(g.Config().Backend.AuthFilter)
	for i, r := range res {
		r.Attributes[usernameAttr][0] = r.Attributes[usernameAttr][0] + domainName
		res[i].Attributes = translateAttributeMap(r.Attributes, g.Config().Backend.AttributesMap)
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
