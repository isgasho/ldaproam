package gocrypto

import (
	"net"
	"testing"
)

const (
	bits = 2048
	host = "ldap.example.org"
	ip   = "127.0.0.1"
)

func Test_genCert(t *testing.T) {
	cert, key, err := GenerateSelfSignedCertKey(host, []net.IP{[]byte(ip)}, []string{"ldap.local"}, bits, 10)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(cert))
	x509 := PemTox509Base64(cert)
	t.Log(string(x509))
	t.Log(string(X509Base64ToPem(x509)))
	t.Log(string(key))
	public, err := GetPublicFromCert(cert)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(public))

}
