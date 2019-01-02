package backend

import (
	"crypto/tls"
	"fmt"

	ldap "gopkg.in/ldap.v2"
)

type LDAP_CONFIG struct {
	Addr       string   `json:"addr"`
	BaseDn     string   `json:"baseDn"`
	BindDn     string   `json:"bindDn`
	BindPass   string   `json:"bindPass"`
	AuthFilter string   `json:"authFilter"`
	Attributes []string `json:"attributes"`
	TLS        bool     `json:"tls"`
	StartTLS   bool     `json:"startTLS"`
	Conn       *ldap.Conn
}

type LDAP_RESULT struct {
	DN         string              `json:"dn"`
	Attributes map[string][]string `json:"attributes"`
}

func (lc *LDAP_CONFIG) Close() {
	if lc.Conn != nil {
		lc.Conn.Close()
		lc.Conn = nil
	}
}

func (lc *LDAP_CONFIG) Connect() (err error) {
	if lc.TLS {
		lc.Conn, err = ldap.DialTLS("tcp", lc.Addr, &tls.Config{InsecureSkipVerify: true})
	} else {
		lc.Conn, err = ldap.Dial("tcp", lc.Addr)
	}
	if err != nil {
		return err
	}
	if !lc.TLS && lc.StartTLS {
		err = lc.Conn.StartTLS(&tls.Config{InsecureSkipVerify: true})
		if err != nil {
			lc.Conn.Close()
			return err
		}
	}

	return err
}

func (lc *LDAP_CONFIG) Bind(BindDn, BindPass string) (err error) {
	err = lc.Conn.Bind(BindDn, BindPass)
	return err
}

func (lc *LDAP_CONFIG) Search(SearchFilter string, attributes []string) (results []LDAP_RESULT, err error) {
	searchRequest := ldap.NewSearchRequest(
		lc.BaseDn, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		SearchFilter, // The filter to apply
		attributes,   // A list attributes to retrieve
		nil,
	)
	sr, err := lc.Conn.Search(searchRequest)
	if err != nil {
		return
	}
	if len(sr.Entries) == 0 {
		return
	}
	results = []LDAP_RESULT{}
	var result LDAP_RESULT
	for _, entry := range sr.Entries {
		attributes := make(map[string][]string)
		for _, attr := range entry.Attributes {
			attributes[attr.Name] = attr.Values
		}
		result.DN = entry.DN
		result.Attributes = attributes
		results = append(results, result)
	}
	return
}

func (lc *LDAP_CONFIG) SearchUser(username string, attributes []string) (results []LDAP_RESULT, err error) {
	searchRequest := ldap.NewSearchRequest(
		lc.BaseDn, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(lc.AuthFilter, username), // The filter to apply
		attributes,                           // A list attributes to retrieve
		nil,
	)
	sr, err := lc.Conn.Search(searchRequest)
	if err != nil {
		return
	}
	if len(sr.Entries) == 0 {
		return
	}
	results = []LDAP_RESULT{}
	var result LDAP_RESULT
	for _, entry := range sr.Entries {
		attributes := make(map[string][]string)
		for _, attr := range entry.Attributes {
			attributes[attr.Name] = attr.Values
		}
		result.DN = entry.DN
		result.Attributes = attributes
		results = append(results, result)
	}
	return
}
