package forward

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/shanghai-edu/ldaproam/backend"
	"github.com/shanghai-edu/ldaproam/g"
	"github.com/shanghai-edu/ldaproam/http/controllers"
	"github.com/shanghai-edu/ldaproam/metadata"
)

func translateAttributes(attributes []string, attributeMap map[string]string) (newAttributes []string) {
	newMap := map[string]string{}
	for k, v := range attributeMap {
		newMap[v] = k
	}

	for _, attr := range attributes {
		if newAttr, ok := newMap[attr]; ok {
			newAttributes = append(newAttributes, newAttr)
		}
	}
	return
}

func translateAttributeMap(LdapResutlAttirbutes map[string][]string, attributeMap map[string]string) map[string][]string {
	newLdapResutlAttirbutes := map[string][]string{}
	for k, v := range LdapResutlAttirbutes {
		if newKey, ok := attributeMap[k]; ok {
			newLdapResutlAttirbutes[newKey] = v
		} else {
			newLdapResutlAttirbutes[k] = v
		}
	}
	return newLdapResutlAttirbutes
}

func BindForward(dn, pass string) (err error) {
	m, err := metadata.GetMetadataByDn(dn)
	if err != nil {
		return
	}
	err, BindReq := CreateBindReq(g.Config().Metadata.DomainName, m.Entity.DomainName, dn, pass, g.Config().Credentials.PrivateKey)
	if err != nil {
		return
	}
	err, body := SendHttpReq([]byte(m.Certificate), m.Endpoint.Bind, BindReq)
	if err != nil {
		return
	}
	var js controllers.MsgResult
	err = json.Unmarshal(body, &js)
	if err != nil {
		return
	}
	if js.Success == false {
		err = errors.New(js.Msg)
		return
	}
	return
}

func SearchForward(username string, domain string, attributes []string) (results []backend.LDAP_RESULT, err error) {
	m, err := metadata.GetMetadataByDomain(domain)
	if err != nil {
		return
	}
	err, SearchReq := CreateSearchReq(g.Config().Metadata.DomainName, m.Entity.DomainName, username, domain, translateAttributes(attributes, g.Config().Backend.AttributesMap), g.Config().Credentials.PrivateKey)
	if err != nil {
		return
	}
	err, body := SendHttpReq([]byte(m.Certificate), m.Endpoint.Search, SearchReq)
	if err != nil {
		return
	}
	g.DebugLog(fmt.Sprintf("Forward Search Response: %s", string(body)))
	var js controllers.SearchResult
	err = json.Unmarshal(body, &js)
	if err != nil {
		return
	}
	if js.Success == false {
		var jss controllers.MsgResult
		json.Unmarshal(body, &jss)
		err = errors.New(jss.Msg)
		return
	}

	for i, LdapResult := range js.Result {
		newAttributes := translateAttributeMap(LdapResult.Attributes, g.Config().Backend.AttributesMap)
		js.Result[i].Attributes = newAttributes
	}
	results = js.Result
	g.DebugLog(fmt.Sprintf("Forward Search Result: %s", js.Result))
	return
}
