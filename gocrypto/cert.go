package gocrypto

import (
	"bytes"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"net"
	"strings"
	"time"

	"github.com/toolkits/file"
)

func GenerateSelfSignedCertKey(host string, alternateIPs []net.IP, alternateDNS []string, bits, after int) ([]byte, []byte, error) {
	priv, err := rsa.GenerateKey(cryptorand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: fmt.Sprintf("%s@%d", host, time.Now().Unix()),
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24 * 365 * time.Duration(after)),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA: true,
	}

	if ip := net.ParseIP(host); ip != nil {
		template.IPAddresses = append(template.IPAddresses, ip)
	} else {
		template.DNSNames = append(template.DNSNames, host)
	}

	derBytes, err := x509.CreateCertificate(cryptorand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return nil, nil, err
	}

	// Generate cert
	certBuffer := bytes.Buffer{}
	if err := pem.Encode(&certBuffer, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return nil, nil, err
	}

	// Generate key
	keyBuffer := bytes.Buffer{}
	if err := pem.Encode(&keyBuffer, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)}); err != nil {
		return nil, nil, err
	}

	return certBuffer.Bytes(), keyBuffer.Bytes(), nil
}

func GetPublicFromCert(cert []byte) ([]byte, error) {
	block, _ := pem.Decode(cert)
	if block == nil {
		return nil, errors.New("cert error!")
	}
	certificate, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPublicKey := certificate.PublicKey.(*rsa.PublicKey)

	keyBuffer := bytes.Buffer{}
	if err := pem.Encode(&keyBuffer, &pem.Block{Type: "RSA PUBLIC KEY", Bytes: x509.MarshalPKCS1PublicKey(rsaPublicKey)}); err != nil {
		return nil, err
	}
	return keyBuffer.Bytes(), nil
}

func PemTox509Base64(cert []byte) []byte {
	pemCert := string(cert)
	x509 := ""
	for _, line := range strings.Split(pemCert, "\n") {
		if strings.Contains(line, "-----BEGIN") {
			continue
		}
		if strings.Contains(line, "-----END") {
			continue
		}
		if line == "" {
			continue
		}
		x509 = x509 + line + "\n"
	}
	return []byte(x509)
}

func X509Base64ToPem(cert []byte) []byte {
	certBuffer := bytes.Buffer{}
	certBuffer.WriteString("-----BEGIN CERTIFICATE-----\n")
	certBuffer.Write(cert)
	certBuffer.WriteString("-----END CERTIFICATE-----\n")
	return certBuffer.Bytes()
}

func GenCert(cerFile, keyFile, host string, alternateIPs []net.IP, alternateDNS []string, bits, after int) error {
	cert, key, err := GenerateSelfSignedCertKey(host, alternateIPs, alternateDNS, bits, after)
	if err != nil {
		return err
	}
	_, err = file.WriteBytes(cerFile, cert)
	if err != nil {
		return err
	}
	_, err = file.WriteBytes(keyFile, key)
	if err != nil {
		return err
	}
	return nil
}
