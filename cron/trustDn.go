package cron

import (
	"log"
	"sync"
	"time"

	"github.com/shanghai-edu/ldaproam/g"
)

var TrustDN sync.Map

func UpdateTrustDN() {
	go updateTrustDN(g.Config().Ldap.TrustDnCache.Interval)
}

func updateTrustDN(sec int64) {
	t := time.NewTicker(time.Second * time.Duration(sec))
	defer t.Stop()
	for {
		TrustDN.Range(func(k, v interface{}) bool {
			dn, valid := k.(string)
			if !valid {
				log.Println("invalid type assertion error", k)
				return true
			}
			expired, valid := v.(time.Time)
			if !valid {
				log.Println("invalid type assertion error", v)
				return true
			}
			now := time.Now()
			if now.After(expired) {
				TrustDN.Delete(dn)
			}
			return true
		})
		log.Println("TrustDN Updated")
	}
}

func AddTrustDN(dn string) error {
	now := time.Now()
	h, err := time.ParseDuration(g.Config().Ldap.TrustDnCache.Expired)
	if err != nil {
		return err
	}
	expired := now.Add(h)
	TrustDN.Store(dn, expired)
	return nil
}
