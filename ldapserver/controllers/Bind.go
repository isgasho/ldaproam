package controllers

import (
	"log"
	"strings"

	"github.com/shanghai-edu/ldaproam/backend"
	"github.com/shanghai-edu/ldaproam/forward"
	"github.com/shanghai-edu/ldaproam/g"
	ldap "github.com/vjeantet/ldapserver"
)

func HandleBind(w ldap.ResponseWriter, m *ldap.Message) {
	r := m.GetBindRequest()
	res := ldap.NewBindResponse(ldap.LDAPResultSuccess)
	if r.AuthenticationChoice() == "simple" {
		dn := string(r.Name())
		pass := string(r.AuthenticationSimple())
		if dn == g.Config().Ldap.BindDn && pass == g.Config().Ldap.BindPass {
			w.Write(res)
			return
		}

		var err error
		if strings.Contains(strings.ToLower(dn), strings.ToLower(g.Config().Metadata.BaseDn)) {
			err = backend.LDAP_Bind(dn, pass)
		} else {
			err = forward.BindForward(dn, pass)
		}

		if err == nil {
			w.Write(res)
			return
		}
		log.Printf("Bind failed User=%s, Pass=%#v", string(r.Name()), r.Authentication(), err)
		res.SetResultCode(ldap.LDAPResultInvalidCredentials)
		res.SetDiagnosticMessage("invalid credentials")
	} else {
		res.SetResultCode(ldap.LDAPResultUnwillingToPerform)
		res.SetDiagnosticMessage("Authentication choice not supported")
	}

	w.Write(res)
}
