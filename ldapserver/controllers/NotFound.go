package controllers

import (
	"log"

	ldap "github.com/vjeantet/ldapserver"
)

func HandleNotFound(w ldap.ResponseWriter, r *ldap.Message) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("ldap client panic recovered in notfound: %v", err)
		}
	}()

	switch r.ProtocolOpType() {
	case ldap.ApplicationBindRequest:
		res := ldap.NewBindResponse(ldap.LDAPResultSuccess)
		res.SetDiagnosticMessage("Default binding behavior set to return Success")

		w.Write(res)

	default:
		res := ldap.NewResponse(ldap.LDAPResultUnwillingToPerform)
		res.SetDiagnosticMessage("Operation not implemented by server")
		w.Write(res)
	}
}
