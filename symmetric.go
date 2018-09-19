package main

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"encoding/hex"
	"errors"
	mrand "math/rand"
	"strconv"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("WhisperGo", map[string]interface{}{
		"EncryptSymmetric": EncryptSymmetric,
	})
}

// from whisperv6/doc.go
const (
	aesKeyLength   = 32 // in bytes
	aesNonceLength = 12 // in bytes; for more info please see cipher.gcmStandardNonceSize & aesgcm.NonceSize()
)

// EncryptSymmetric encrypts a message with a topic key, using AES-GCM-256.
// nonce size should be 12 bytes (see cipher.gcmStandardNonceSize).
func EncryptSymmetric(key []byte, msg string) (string, error) {
	if !validateDataIntegrity(key, aesKeyLength) {
		return "", errors.New("invalid key provided for symmetric encryption, size: " + strconv.Itoa(len(key)))
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	salt, err := generateSecureRandomData(aesNonceLength) // never use more than 2^32 random nonces with a given key
	if err != nil {
		return "", err
	}

	var raw []byte
	encrypted := aesgcm.Seal(nil, salt, []byte(msg), nil)
	raw = append(encrypted, salt...)

	return hex.EncodeToString(raw), nil
}

// validateDataIntegrity returns false if the data have the wrong or contains all zeros,
// which is the simplest and the most common bug.
func validateDataIntegrity(k []byte, expectedSize int) bool {
	if len(k) != expectedSize {
		return false
	}
	if expectedSize > 3 && containsOnlyZeros(k) {
		return false
	}
	return true
}

// containsOnlyZeros checks if the data contain only zeros.
func containsOnlyZeros(data []byte) bool {
	for _, b := range data {
		if b != 0 {
			return false
		}
	}
	return true
}

// generateSecureRandomData generates random data where extra security is required.
// The purpose of this function is to prevent some bugs in software or in hardware
// from delivering not-very-random data. This is especially useful for AES nonce,
// where true randomness does not really matter, but it is very important to have
// a unique nonce for every message.
func generateSecureRandomData(length int) ([]byte, error) {
	x := make([]byte, length)
	y := make([]byte, length)
	res := make([]byte, length)

	_, err := crand.Read(x)
	if err != nil {
		return nil, err
	} else if !validateDataIntegrity(x, length) {
		return nil, errors.New("crypto/rand failed to generate secure random data")
	}
	_, err = mrand.Read(y)
	if err != nil {
		return nil, err
	} else if !validateDataIntegrity(y, length) {
		return nil, errors.New("math/rand failed to generate secure random data")
	}
	for i := 0; i < length; i++ {
		res[i] = x[i] ^ y[i]
	}
	if !validateDataIntegrity(res, length) {
		return nil, errors.New("failed to generate secure random data")
	}
	return res, nil
}
