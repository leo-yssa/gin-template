package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func GenerateKeyPair() ([]byte, []byte, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	b, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return nil, nil, err
	}
	privBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: b,
	}
	b, err = x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return nil, nil, err
	}
	pubBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: b,
	}
	return pem.EncodeToMemory(pubBlock), pem.EncodeToMemory(privBlock), nil
}

func ExtractPrivateKey(data interface{}) (ed25519.PrivateKey, error) {
	pemBytes, ok := data.([]byte)
	if !ok {
		return nil, errors.New("Invalid type")
	}
	p, _ := pem.Decode(pemBytes)
	key, err := x509.ParsePKCS8PrivateKey(p.Bytes)
	if err != nil {
		return nil, err
	}
	ed25519Key, ok := key.(ed25519.PrivateKey)
	if !ok {
		return nil, errors.New("Invalid type")
	}
	return ed25519Key, nil
}