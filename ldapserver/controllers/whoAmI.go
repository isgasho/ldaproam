package controllers

import (
	ldap "github.com/vjeantet/ldapserver"
)

func HandleWhoAmI(w ldap.ResponseWriter, m *ldap.Message) {
	res := ldap.NewExtendedResponse(ldap.LDAPResultSuccess)
	w.Write(res)
}
