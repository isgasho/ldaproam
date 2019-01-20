package forward

import (
	"testing"

	"github.com/shanghai-edu/ldaproam/g"
	"github.com/shanghai-edu/ldaproam/metadata"
)

func init() {
	g.ParseConfig(`D:\GOPATH\src\github.com\shanghai-edu\ldaproam\cfg-bob.json`)
	g.PasseCredentials()
	metadata.InitMetadata()
}

const (
	bindurl   = "https://ldap.a.example.org:8443/api/v1/bind"
	searchurl = "https://ldap.a.example.org:8443/api/v1/search"
)

func Test_bindForward(t *testing.T) {
	res, err := SearchForward("001", "a.example.org", []string{"uid", "cn", "mail"})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
	dn := res[0].DN
	err = BindForward(dn, "123456")
	if err != nil {
		t.Error(err)
	}
	t.Log(err)
}

func Test_GetUsernameFromFilter(t *testing.T) {
	filter := "(&(uid=username))"
	username := GetUsernameFromFilter(filter)
	t.Log(username)
}

func Test_TranslateAttributes(t *testing.T) {
	attributes := []string{"samAccountName", "displayName", "email"}
	attributeMap := map[string]string{
		"uid":  "samAccountName",
		"cn":   "displayName",
		"mail": "email",
	}
	t.Log(translateAttributes(attributes, attributeMap))
}
