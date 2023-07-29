package util

import (
	"encoding/hex"

	ecies "github.com/ecies/go/v2"
)

func KeyGen() (string, string, error) {
	sk, err := ecies.GenerateKey()
	if err != nil {
		return "", "", err
	}
	return sk.PublicKey.Hex(true), sk.Hex(), nil
}

func Encypt(pk string, msg string) (string, error) {
	publickey, err := ecies.NewPublicKeyFromHex(pk)
	if err != nil {
		return "", err
	}
	ct, err := ecies.Encrypt(publickey, []byte(msg))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(ct), nil
}

func Decypt(sk string, ct string) (string, error) {
	privatekey, err := ecies.NewPrivateKeyFromHex(sk)
	if err != nil {
		return "", err
	}
	ctByte, err := hex.DecodeString(ct)
	if err != nil {
		return "", err
	}
	msg, err := ecies.Decrypt(privatekey, ctByte)
	if err != nil {
		return "", err
	}
	return string(msg), nil
}
