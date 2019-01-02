package backend

import (
	"github.com/shanghai-edu/ldaproam/g"
)

func LDAP_Search(SearchFilter string, attributes []string) (results []LDAP_RESULT, err error) {
	lc := &LDAP_CONFIG{
		Addr:       g.Config().Backend.Addr,
		BaseDn:     g.Config().Backend.BaseDn,
		BindDn:     g.Config().Backend.BindDn,
		BindPass:   g.Config().Backend.BindPass,
		AuthFilter: g.Config().Backend.AuthFilter,
		Attributes: g.Config().Backend.Attributes,
		TLS:        g.Config().Backend.TLS,
		StartTLS:   g.Config().Backend.StartTLS,
		Conn:       nil,
	}
	lc.Connect()
	defer lc.Close()
	err = lc.Bind(lc.BindDn, lc.BindPass)
	if err != nil {
		return
	}
	results, err = lc.Search(SearchFilter, attributes)
	return
}

func LDAP_SearchUser(username string, attributes []string) (results []LDAP_RESULT, err error) {
	lc := &LDAP_CONFIG{
		Addr:       g.Config().Backend.Addr,
		BaseDn:     g.Config().Backend.BaseDn,
		BindDn:     g.Config().Backend.BindDn,
		BindPass:   g.Config().Backend.BindPass,
		AuthFilter: g.Config().Backend.AuthFilter,
		Attributes: g.Config().Backend.Attributes,
		TLS:        g.Config().Backend.TLS,
		StartTLS:   g.Config().Backend.StartTLS,
		Conn:       nil,
	}
	lc.Connect()
	defer lc.Close()
	err = lc.Bind(lc.BindDn, lc.BindPass)
	if err != nil {
		return
	}
	results, err = lc.SearchUser(username, attributes)
	return
}

func LDAP_Bind(dn, pass string) (err error) {
	lc := &LDAP_CONFIG{
		Addr:       g.Config().Backend.Addr,
		BaseDn:     g.Config().Backend.BaseDn,
		BindDn:     g.Config().Backend.BindDn,
		BindPass:   g.Config().Backend.BindPass,
		AuthFilter: g.Config().Backend.AuthFilter,
		Attributes: g.Config().Backend.Attributes,
		TLS:        g.Config().Backend.TLS,
		StartTLS:   g.Config().Backend.StartTLS,
		Conn:       nil,
	}
	lc.Connect()
	defer lc.Close()
	err = lc.Bind(dn, pass)
	return
}
