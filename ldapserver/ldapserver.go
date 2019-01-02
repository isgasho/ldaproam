package ldapserver

import (
	"github.com/shanghai-edu/ldaproam/g"
	"github.com/shanghai-edu/ldaproam/ldapserver/controllers"

	ldap "github.com/vjeantet/ldapserver"
)

func Start() {
	//Create a new LDAP Server
	server := ldap.NewServer()

	//Create routes bindings
	routes := ldap.NewRouteMux()
	routes.NotFound(controllers.HandleNotFound)
	routes.Abandon(controllers.HandleAbandon)
	routes.Bind(controllers.HandleBind)

	routes.Extended(controllers.HandleStartTLS).
		RequestName(ldap.NoticeOfStartTLS).Label("StartTLS")

	routes.Extended(controllers.HandleWhoAmI).
		RequestName(ldap.NoticeOfWhoAmI).Label("Ext - WhoAmI")

	routes.Search(controllers.HandleSearch).Label("Search - Generic")

	//Attach routes to server
	server.Handle(routes)

	// listen and serve
	addr := g.Config().Ldap.Listen
	server.ListenAndServe(addr)

}
