package forward

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/shanghai-edu/ldaproam/g"
	"github.com/shanghai-edu/ldaproam/gocrypto"
)

func CreateBindReq(from, to, dn, password string, privateKey []byte) (error, []byte) {
	var HttpReq g.HttpBindReq

	HttpReq.Body.From = from
	HttpReq.Body.To = to
	HttpReq.Body.Data.Dn = dn
	HttpReq.Body.Data.Password = password

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

func SendHttpReq(cert []byte, url string, data []byte) (error, []byte) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err, nil
	}
	req.Header.Set("Content-Type", "application/json")

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(cert)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: pool},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}

	return nil, body
}
