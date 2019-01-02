package cron

import (
	"log"
	"time"

	"github.com/shanghai-edu/ldaproam/g"
	"github.com/shanghai-edu/ldaproam/metadata"
)

func UpdateMetadata() {
	go updatetMetadata(g.Config().Metadata.Interval)
}

func updatetMetadata(sec int64) {
	t := time.NewTicker(time.Second * time.Duration(sec))
	defer t.Stop()
	for {
		<-t.C
		err := metadata.GetMetadata(g.Config().Metadata.Provider, g.Config().Metadata.BackingFile)
		if err != nil {
			log.Println(err)
			continue
		}
		err = metadata.LoadMetadata(g.Config().Metadata.BackingFile)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("Metadata Updated")

	}
}
