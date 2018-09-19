package main

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

func Key(password string) string {
	return string(pbkdf2.Key([]byte(password), nil, 65356, aesKeyLength, sha256.New))
}

// KeyEncrypt is a wrapper around Key + EncryptSymmetric.
func KeyEncrypt(password, msg string) (out string, err string) {
	key := Key(password)
	return EncryptSymmetric(key, msg)
}
