package main

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

func Key(password []byte) []byte {
	return pbkdf2.Key(password, nil, 65356, aesKeyLength, sha256.New)
}
