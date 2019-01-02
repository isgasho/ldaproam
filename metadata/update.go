package metadata

import (
	"bytes"
	"encoding/json"

	"log"
	"net/http"

	"github.com/shanghai-edu/ldaproam/g"
	"github.com/toolkits/file"
)

func InitMetadata() {
	err := GetMetadata(g.Config().Metadata.Provider, g.Config().Metadata.BackingFile)
	if err != nil {
		log.Println(err)
		return
	}
	err = LoadMetadata(g.Config().Metadata.BackingFile)
	if err != nil {
		log.Println(err)
		return
	}
}

func LoadMetadata(metadataFile string) error {
	var metadatas []g.MetadataLoad
	metadataStr, err := file.ToTrimString(metadataFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(metadataStr), &metadatas)
	if err != nil {
		return err
	}
	g.Config().Metadata.Metadatas = metadatas
	return nil
}

func GetMetadata(url, backingFile string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	metadataString := buf.String()

	_, err = file.WriteString(backingFile, metadataString)
	return
}
