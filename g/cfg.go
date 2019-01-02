package g

import (
	"encoding/json"

	"io/ioutil"
	"log"
	"sync"

	"github.com/shanghai-edu/ldaproam/gocrypto"
	"github.com/toolkits/file"
)

type GlobalConfig struct {
	Backend     *BackendConfig     `json:"backend"`
	Credentials *CredentialsConfig `json:"credentials"`
	Metadata    *MetadataConfig    `json:"metadata"`
	Ldap        *LdapServerConfig  `json:"ldapServer"`
	Http        *HttpConfig        `json:"http"`
}

type BackendConfig struct {
	Addr          string            `json:"addr"`
	BaseDn        string            `json:"baseDn"`
	BindDn        string            `json:"bindDn`
	BindPass      string            `json:"bindPass"`
	AuthFilter    string            `json:"authFilter"`
	Attributes    []string          `json:"attributes"`
	AttributesMap map[string]string `json:"attributesMap"`
	TLS           bool              `json:"tls"`
	StartTLS      bool              `json:"startTLS"`
}
type CredentialsConfig struct {
	Bits        int    `json:"bits"`
	Expired     int    `json:"expired"`
	Cert        string `json:"cert"`
	Key         string `json:"key"`
	PrivateKey  []byte `json:"privateKey"`
	PublicKey   []byte `json:"publicKey"`
	Certificate []byte `json:"Certificate"`
}

type MetadataConfig struct {
	DomainName    string         `json:"domainName"`
	ServedDomains []string       `json:"servedDomains"`
	BaseDn        string         `json:"baseDn"`
	Description   string         `json:"description"`
	ApiAddr       string         `json:"apiAddr"`
	Interval      int64          `json:"interval"`
	Provider      string         `json:"provider"`
	BackingFile   string         `json:backingFile"`
	Metadatas     []MetadataLoad `json:"metadatas"`
}

type MetadataLoad struct {
	Entity      EntityConfig   `json:"entity"`
	Certificate string         `json:"certificate"`
	Endpoint    EndpointConfig `json:"endpoint"`
}

type EntityConfig struct {
	DomainName    string   `json:"domainName`
	ServedDomains []string `json:"servedDomains"`
	BaseDn        string   `json:"baseDn"`
	Description   string   `json:"description"`
}

type EndpointConfig struct {
	Bind   string `json:"bind"`
	Search string `json:"search"`
}
type LdapServerConfig struct {
	Debug        bool        `json:"debug"`
	BindDn       string      `json:"bindDn`
	BindPass     string      `json:"bindPass"`
	TrustDnCache *TrustCache `json:"trustDnCache"`
	Listen       string      `json:"listen"`
}

type TrustCache struct {
	Expired  string `json:"expired"`
	Interval int64  `json:"interval"`
}

type HttpConfig struct {
	Debug  bool   `json:"debug"`
	Listen string `json:"listen"`
	Port   int    `json:"port"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	lock       = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent. maybe you need `mv cfg.example.json cfg.json`")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	lock.Lock()
	defer lock.Unlock()

	config = &c
}

func PasseCredentials() {
	var err error
	Config().Credentials.Certificate, err = ioutil.ReadFile(Config().Credentials.Cert)
	if err != nil {
		log.Fatalln("parse Credentials Cert file:", Config().Credentials.Cert, "fail:", err)
	}
	Config().Credentials.PrivateKey, err = ioutil.ReadFile(Config().Credentials.Key)
	if err != nil {
		log.Fatalln("parse Credentials Key file:", Config().Credentials.Key, "fail:", err)
	}
	PublicKey, err := gocrypto.GetPublicFromCert(Config().Credentials.Certificate)
	if err != nil {
		log.Fatalln("parse Public Key fail:", err)
	}
	Config().Credentials.PublicKey = PublicKey
}
