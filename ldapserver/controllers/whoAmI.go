package controllers

import (
	"log"

	ldap "github.com/vjeantet/ldapserver"
)

func HandleWhoAmI(w ldap.ResponseWriter, m *ldap.Message) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("ldap client panic recovered in who am i: %v", err)
		}
	}()

	res := ldap.NewExtendedResponse(ldap.LDAPResultSuccess)
	w.Write(res)
}
