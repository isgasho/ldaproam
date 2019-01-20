package controllers

import (
	"fmt"
	"log"
	"strings"

	"github.com/shanghai-edu/ldaproam/forward"

	"github.com/shanghai-edu/ldaproam/backend"
	"github.com/shanghai-edu/ldaproam/g"
	"github.com/shanghai-edu/ldaproam/metadata"
	"github.com/vjeantet/goldap/message"
	ldap "github.com/vjeantet/ldapserver"
)

func HandleSearch(w ldap.ResponseWriter, m *ldap.Message) {
	r := m.GetSearchRequest()

	// Handle Stop Signal (server stop / client disconnected / Abandoned request....)
	select {
	case <-m.Done:
		log.Print("Leaving handleSearch...")
		return
	default:
	}
	attrs := r.Attributes()
	attributes := []string{}
	for _, attr := range attrs {
		attributes = append(attributes, string(attr))
	}
	var results []backend.LDAP_RESULT
	var err error

	username := forward.GetUsernameFromFilter(r.FilterString())
	usernameSplit := strings.Split(username, "@")

	g.DebugLog(fmt.Sprintf("FilterString: %s", r.FilterString()))

	if len(usernameSplit) == 2 {
		g.DebugLog(fmt.Sprintf("Username: %s , Domain: %s", usernameSplit[0], usernameSplit[1]))
		if metadata.InArray(usernameSplit[1], g.Config().Metadata.ServedDomains) {
			g.DebugLog("Search Location: Seach Backend")
			results, err = backend.LDAP_Search(r.FilterString(), attributes)
		} else {
			g.DebugLog("Search Location: Search Forward")
			results, err = forward.SearchForward(usernameSplit[0], usernameSplit[1], attributes)
		}
	} else {
		g.DebugLog("Search Location: Search Backend")
		results, err = backend.LDAP_Search(r.FilterString(), attributes)
	}

	if err != nil {
		log.Println(err)
		serverRes := ldap.NewSearchResultDoneResponse(ldap.LDAPResultUnavailable)
		w.Write(serverRes)
	}

	for _, res := range results {
		e := ldap.NewSearchResultEntry(res.DN)
		for name, values := range res.Attributes {
			for _, value := range values {
				e.AddAttribute(message.AttributeDescription(name), message.AttributeValue(value))
			}
		}
		w.Write(e)
	}
	serverRes := ldap.NewSearchResultDoneResponse(ldap.LDAPResultSuccess)
	w.Write(serverRes)

}
