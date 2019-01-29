package forward

import (
	"strings"

	"encoding/json"

	"github.com/shanghai-edu/ldaproam/g"
	"github.com/shanghai-edu/ldaproam/gocrypto"
)

func GetUsernameFromFilter(filter string) (username string) {
	split := strings.Split(filter, "=")
	split2 := strings.Split(split[1], ")")
	username = split2[0]
	return
}

func CreateSearchReq(from, to, username, domain string, attributes []string, privateKey []byte) (error, []byte) {
	var HttpReq g.HttpSearchReq

	HttpReq.Body.From = from
	HttpReq.Body.To = to
	HttpReq.Body.Data.Username = username
	HttpReq.Body.Data.Domain = domain
	HttpReq.Body.Data.Attributes = attributes

	b, err := json.Marshal(HttpReq.Body)
	if err != nil {
		return err, nil
	}
	sign, err := gocrypto.RsaSign(string(b), privateKey)
	if err != nil {
		return err, nil
	}
	HttpReq.Sign = sign
	b, err = json.Marshal(HttpReq)
	if err != nil {
		return err, nil
	}
	return err, b
}
