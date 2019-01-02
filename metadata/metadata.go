package metadata

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/shanghai-edu/ldaproam/g"
	"github.com/toolkits/file"
)

func GenMetadata(metadata_file string) error {
	endpoint := g.EndpointConfig{
		Bind:   fmt.Sprintf("%s/bind", g.Config().Metadata.ApiAddr),
		Search: fmt.Sprintf("%s/search", g.Config().Metadata.ApiAddr),
	}
	entity := g.EntityConfig{
		DomainName:    g.Config().Metadata.DomainName,
		ServedDomains: g.Config().Metadata.ServedDomains,
		BaseDn:        g.Config().Metadata.BaseDn,
		Description:   g.Config().Metadata.Description,
	}
	metadata := g.MetadataLoad{
		Entity:      entity,
		Certificate: string(g.Config().Credentials.Certificate),
		Endpoint:    endpoint,
	}
	js_data, err := json.MarshalIndent(metadata, "", "\t")
	if err != nil {
		return err
	}
	_, err = file.WriteBytes(metadata_file, js_data)
	if err != nil {
		return err
	}
	return nil
}

func GetMetadataByDomainName(domainName string) (metadata g.MetadataLoad, err error) {
	for _, m := range g.Config().Metadata.Metadatas {
		if m.Entity.DomainName == domainName {
			metadata = m
			return
		}
	}
	err = errors.New("Cannot Found Such DomainName")
	return
}

func GetMetadataByDn(dn string) (metadata g.MetadataLoad, err error) {
	for _, m := range g.Config().Metadata.Metadatas {
		if strings.Contains(strings.ToLower(dn), strings.ToLower(m.Entity.BaseDn)) {
			metadata = m
			return
		}
	}
	err = errors.New("Cannot Found Such DN")
	return
}

func GetMetadataByDomain(domain string) (metadata g.MetadataLoad, err error) {
	for _, m := range g.Config().Metadata.Metadatas {
		if InArray(domain, m.Entity.ServedDomains) {
			metadata = m
			return
		}
	}
	err = errors.New("Cannot Found Such Domain")
	return
}

func InArray(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
