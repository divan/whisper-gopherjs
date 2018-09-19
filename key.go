package main

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

func Key(password, salt []byte, iter, keyLen int) []byte {
	return pbkdf2.Key(password, salt, iter, keyLen, sha256.New)
}
