package gocrypto

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"os"
)

func GenRsaKey(bits int, private, public string) error {
	// gen private key
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "private",
		Bytes: derStream,
	}
	file, err := os.Create(private)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// gen public key
	publicKey := &privateKey.PublicKey
	derPkix := x509.MarshalPKCS1PublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "public",
		Bytes: derPkix,
	}
	file, err = os.Create(public)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

func RsaEncrypt(data string, publicKey []byte) (string, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return "", errors.New("public key error")
	}
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	partLen := pub.N.BitLen()/8 - 11
	chunks := split([]byte(data), partLen)
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		bytes, err := rsa.EncryptPKCS1v15(rand.Reader, pub, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(bytes)
	}

	return base64.StdEncoding.EncodeToString(buffer.Bytes()), err
}

func RsaDecrypt(ciphertext string, privateKey []byte) (string, error) {

	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	partLen := priv.N.BitLen() / 8
	raw, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	chunks := split([]byte(raw), partLen)
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		decryptData, err := rsa.DecryptPKCS1v15(rand.Reader, priv, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(decryptData)
	}
	return buffer.String(), err
}

func RsaSign(data string, privateKey []byte) (string, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	hash := sha256.New()
	hash.Write([]byte(data))
	signature, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hash.Sum(nil))
	return base64.StdEncoding.EncodeToString(signature), err
}

func RsaVerify(data, signature string, publicKey []byte) error {
	sign, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("public key error")
	}
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return err
	}
	hash := sha256.New()
	hash.Write([]byte(data))
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hash.Sum(nil), sign)
}

func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}
