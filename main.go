package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/shanghai-edu/ldaproam/cron"
	"github.com/shanghai-edu/ldaproam/g"
	"github.com/shanghai-edu/ldaproam/gocrypto"
	"github.com/shanghai-edu/ldaproam/http"
	"github.com/shanghai-edu/ldaproam/ldapserver"
	"github.com/shanghai-edu/ldaproam/metadata"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	gen := flag.Bool("gen", false, "gen certificate")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)
	if *gen {
		err := gocrypto.GenCert(g.Config().Credentials.Cert,
			g.Config().Credentials.Key,
			g.Config().Metadata.DomainName,
			[]net.IP{},
			[]string{},
			g.Config().Credentials.Bits,
			g.Config().Credentials.Expired)
		g.PasseCredentials()
		err = metadata.GenMetadata("./metadata.json")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("gen Credentials Success: ", g.Config().Credentials.Cert, g.Config().Credentials.Key)
			fmt.Println("gen Metadata Success, metadata.json")
		}
		os.Exit(0)
	}
	g.PasseCredentials()

	metadata.InitMetadata()

	go cron.UpdateMetadata()
	go ldapserver.Start()
	go http.Start()

	select {}
}
