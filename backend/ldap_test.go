package backend

import (
	"testing"

	"github.com/shanghai-edu/ldaproam/g"
)

func init() {
	g.ParseConfig(`D:\GOPATH\src\github.com\shanghai-edu\ldaproam\cfg-bob.json`)
}

func Test_LDAP_func(t *testing.T) {
	res, err := LDAP_Search("(displayName=bob)", []string{"uid", "cn", "mail"})
	t.Log(res, err)
	res, err = LDAP_SearchUser("002", []string{"uid", "cn", "mail"})
	t.Log(res, err)
	err = LDAP_Bind("Administrator@b.example.org", "password")
	t.Log(err)
}
