package gocrypto

import (
	"testing"
)

const (
	certificate = `-----BEGIN CERTIFICATE-----
MIIDIjCCAgqgAwIBAgIBATANBgkqhkiG9w0BAQsFADAoMSYwJAYDVQQDDB1sZGFw
LmIuZXhhbXBsZS5vcmdAMTU0NDEwNzQxNTAeFw0xODEyMDYxNDQzMzVaFw0yODEy
MDMxNDQzMzVaMCgxJjAkBgNVBAMMHWxkYXAuYi5leGFtcGxlLm9yZ0AxNTQ0MTA3
NDE1MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0C0QY0xJYfBi7u8R
HVRWlxDB8fviIEro5Zl5suPX3sogG0ncdyY8K53QbTqxki4V5mF3BSZqxt5YKJnX
RDHTmdLdy1qyrmn+GmBnWlU6jhTvCbZbtbjsyTsfb2AdOc88z/nck2ERI+ndtzlg
hpNLPZiA+EsLIL7Mjb8rYVBR6tE6Rvqbz8zqyibdcEd4NEYcjFQnVh2iTa+uDIPs
MNzOwTnWihMR3SIWAfmoxnbsvt9Z15xeNedy3C3YCUSREGAalsLqsYlSpNDNw9i/
2rtwnZlB/KA9z3cAg//SSvrmhuEgJq77i0IVFWv3VgGWmJxZv6ER54FRCnzn5fNT
THmRvQIDAQABo1cwVTAOBgNVHQ8BAf8EBAMCAqQwEwYDVR0lBAwwCgYIKwYBBQUH
AwEwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHREEFjAUghJsZGFwLmIuZXhhbXBsZS5v
cmcwDQYJKoZIhvcNAQELBQADggEBAHx/nQuaSbXAjDY2w7Gt0LvAO0TIkGpox74N
C+5bqXxfQDUUEEFNzF2YJ6tV7KV+3pjIj0tkgP6XyLV202F/Nd4BiWAWkYajpGfV
/NAYIY7alDi9KH08N+IHQwfPjgcWic3BXTv7DbaZUqkPteVxnBVRgh8qLsVXgd3D
8K8cB/52Ck7yU/RiUcLR2+Rn3XIVban1085d5mMZxLDkkMfbo0DUzfZzj/Em1LfA
0ydW22fvqdI1rmf1/xLE8rkg9brufWhjzUu3tbO8nZfB7Nl33/eRg/UcJel318eH
Tjf3R6ugGbdNHPG4hc6rUCiCXuEXiEsh+ZsW4EOfOcAtS3IMIEg=
-----END CERTIFICATE-----

`
	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA0C0QY0xJYfBi7u8RHVRWlxDB8fviIEro5Zl5suPX3sogG0nc
dyY8K53QbTqxki4V5mF3BSZqxt5YKJnXRDHTmdLdy1qyrmn+GmBnWlU6jhTvCbZb
tbjsyTsfb2AdOc88z/nck2ERI+ndtzlghpNLPZiA+EsLIL7Mjb8rYVBR6tE6Rvqb
z8zqyibdcEd4NEYcjFQnVh2iTa+uDIPsMNzOwTnWihMR3SIWAfmoxnbsvt9Z15xe
Nedy3C3YCUSREGAalsLqsYlSpNDNw9i/2rtwnZlB/KA9z3cAg//SSvrmhuEgJq77
i0IVFWv3VgGWmJxZv6ER54FRCnzn5fNTTHmRvQIDAQABAoIBAQChepM/sykpA2J6
nI7WBVm1jJYkspHRIspNbyNrlDka5AbLpBuTgOEjpCLM175t+bmbHvdH02j4IfiY
Zd0JVO0iYOMSnqQDjsxAgY8qDvqAw9Q34HB7IZrq7SWRKykAcrRlTxe+aoj1Jq+J
NMSfHxo6CVXhQ7S6DcZ2HVf1AGzKi2IYZh+RETSpWcwYhR1Nbsg9U2PrbsWlsW7J
guYykG3g99pEm2Y5WHNmZIrznR0ttHq+W3Z1tLa5L45hSCzlHOiS6nxJ3s5EBqUu
da3Ax4OXZermTS3Xhhzi2Iof8oq2WGeuod5LhWdjRfl3X47MPT56gXdburWwgBK6
R+PcM6AxAoGBAPr3pGJskMwOF/mRoygJCcQzBaz4Bu5n68wjl+qBxdGftyDYr2pH
UQh/9lV6IDkJml8Pxm9tbAIolc3xzhbcdP0p3CbX7cO7PoT4yHNjr4Ja/CqGXP5Q
XRS8uUs3acBMYLspIQGLlAnrW2jNe3ZRI/C9QvMgaPAMCJ4NcIyNtjP7AoGBANRZ
v++WRho1q6MjTc8eYklXdIrZwSECy9ED2EmwQ9hxbhtp7tR15XQ+eR23a2zj+SNg
AajXDMV0X3/nsL2dmcnP1FaH85+oQzdX+Jv1N56URD+gXJE77IQjBUdmvCRFKWfq
gu70CeHJaSSSwGOzEzzBgb6B321JCSFoDuwHn6unAoGAICjHqc0bqOpNbC+bZq/x
znBzU7zctoQelSQifWxvuvLqdo0NvWKyIZK0MDPcGTL/0xqkZPbyljw5JhDMReWu
IBrTGS8mSqSd2FBA73hgryWVlqVtGTGXG+crH1ZUeM3Qv2r/zcDjEXpVVlKudXTk
VB8Mizcl+0yvdgFm4LvwEy0CgYAP9gcFPShbw/j4tCifDsuYc6hg32ky0AD93uoc
79DJrgz2pom7EnmCuUdlQmoirygEzqyRQkjFdq/O711Lg1MR5jsxndpj/8O9nzEi
l9XsZ3yRw73xdK2caP12lnRBzakFFI1u5Izxma/7fcRUOhuSD4FvDlf64Oh8yFOG
zjPkNQKBgAuF4IoRztQKO/1W8NYxzYam1cYziG2wpWDExYmmdeNP77ZH186jln8o
i/F1mtUTON1YW7/UyZzYXVqm6zHUf+/bhT09BclAkvEbsQnxWDS7EIps5kt1iQTw
GkOcQOZZCSexRJU0eZkStS3uB6jUwRTsvOWDTc9sgNIEhE5lF5g8
-----END RSA PRIVATE KEY-----`

	exampleCertificate = "-----BEGIN CERTIFICATE-----\nMIIDIjCCAgqgAwIBAgIBATANBgkqhkiG9w0BAQsFADAoMSYwJAYDVQQDDB1sZGFw\nLmEuZXhhbXBsZS5vcmdAMTU0NDE1MjY4NjAeFw0xODEyMDcwMzE4MDZaFw0yODEy\nMDQwMzE4MDZaMCgxJjAkBgNVBAMMHWxkYXAuYS5leGFtcGxlLm9yZ0AxNTQ0MTUy\nNjg2MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo1314VTTPPR118pM\npZrSU9Ohy+UCTX/hwT0eksRRDSwGZobm0yv0qZE7pqaMe7QRX/E18jsW4rwREd5+\ntPjkFTZv4IV1t/6QELh0F6x027USM38RYza8ffREPeHWKqNbnRGGquNTWYbBQ6Hc\nHQCOd9uWGEWs4PvfKnOSDN/3tUiunRTlRZFRlDSaphGOvrdsxJ048CmcF3OW89H4\nczEPFxI6lHIxrv3hqXNF+hndBhsWV3Sio4vvL6dmHnfzbkg67S7ahMp3xsYmkubz\n3hQ98knaao+lUaABoirDgbz9YvMg2VEyYW6FhUOabUYLH8KRabu+6SFMoVJu5SOM\n06I4SwIDAQABo1cwVTAOBgNVHQ8BAf8EBAMCAqQwEwYDVR0lBAwwCgYIKwYBBQUH\nAwEwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHREEFjAUghJsZGFwLmEuZXhhbXBsZS5v\ncmcwDQYJKoZIhvcNAQELBQADggEBADzX6LZx4yKVGPvuLPAmiOWHVDtSrjmp9DUJ\nP8z4QKC+ErGDYbFQBhS+Sm8j0mv9uU18F3+zJry613Y9y5znOUSnG8HEmwx4fVtC\nCxdT1fkUKXHk5Lio7gcPwOofJ4EGuQXCTrLrRlSag3soFftgNqgyi4Mhtquywy26\na0R2XNY7+1QOvafPNM6BRr7tiTfLn/v43M+YEnTxVOlouwUmVrIxHznHGcDuPWRt\nUzTd3+PmRAytmLt1fpM0xr9TsqtSrFqYYrEH/aVQB798/CuctV0kkF4HPZ047nPb\nseF/yZ/HprdzS9oNjTj1HvDbJk2yAIR81MYzVoPrkcK0142c3qw=\n-----END CERTIFICATE-----\n"
)

func Test_rsa_Sign(t *testing.T) {

	publicKey, err := GetPublicFromCert([]byte(certificate))
	if err != nil {
		t.Error(err)
	}
	data := `{"from":"a.example.org","to":"b.exampe.org","data":{"dn":"uid=username,dc=exmaple,dc=org","password":"123456"}}`
	sign, err := RsaSign(data, []byte(privateKey))
	if err != nil {
		t.Error(err)
	}
	t.Log(sign)
	err = RsaVerify(data, sign, publicKey)
	t.Log(err)
}

func Test_rsa_encrypt(t *testing.T) {
	str := "123456"
	publicKey, err := GetPublicFromCert([]byte(certificate))
	if err != nil {
		t.Error(err)
	}
	encryptStr, err := RsaEncrypt(str, publicKey)
	if err != nil {
		t.Error(err)
	}
	t.Log(encryptStr)
	decryptStr, err := RsaDecrypt(encryptStr, []byte(privateKey))
	if err != nil {
		t.Error(err)
	}
	t.Log(decryptStr)
	examplePublicKey, err := GetPublicFromCert([]byte(exampleCertificate))
	if err != nil {
		t.Error(err)
	}
	exampleEncryptStr, err := RsaEncrypt(str, examplePublicKey)
	if err != nil {
		t.Error(err)
	}
	t.Log(exampleEncryptStr)
}
