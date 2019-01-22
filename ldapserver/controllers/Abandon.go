package controllers

import (
	"log"

	ldap "github.com/vjeantet/ldapserver"
)

func HandleAbandon(w ldap.ResponseWriter, m *ldap.Message) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("ldap client panic recovered in abandon: %v", err)
		}
	}()

	var req = m.GetAbandonRequest()
	// retreive the request to abandon, and send a abort signal to it
	if requestToAbandon, ok := m.Client.GetMessageByID(int(req)); ok {
		requestToAbandon.Abandon()
		log.Printf("Abandon signal sent to request processor [messageID=%d]", int(req))
	}
}
