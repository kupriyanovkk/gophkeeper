package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Crypter interface {
	Encode(payload string) string
	Decode(sha string) (string, error)
}

type crypt struct {
	nonce []byte
	aead  cipher.AEAD
	block cipher.Block
}

var password = "SECRET_PASSWORD"

// Encode description of the Go function.
//
// It takes a payload string as a parameter and returns a string.
func (c *crypt) Encode(payload string) string {
	src := []byte(payload)

	dst := c.aead.Seal(nil, c.nonce, src, nil)

	sha := hex.EncodeToString(dst)

	return sha
}

// Decode decodes the sha using the crypt receiver.
//
// Parameter(s): sha string
// Return type(s): string, error
func (c *crypt) Decode(sha string) (string, error) {
	dst, errDecode := hex.DecodeString(sha)
	if errDecode != nil {
		return "", fmt.Errorf("hex decode error: %w", errDecode)
	}

	src, errGCM := c.aead.Open(nil, c.nonce, dst, nil)
	if errGCM != nil {
		return "", errDecode
	}

	return string(src), nil
}

// NewCrypt creates a new crypt object.
//
// None.
// *crypt, error
func NewCrypt() (*crypt, error) {
	key := sha256.Sum256([]byte(password))

	aesBlock, errBlock := aes.NewCipher(key[:])
	if errBlock != nil {
		return nil, fmt.Errorf("aes.NewCipher error: %w", errBlock)
	}

	aesGCM, errGCM := cipher.NewGCM(aesBlock)
	if errGCM != nil {
		return nil, fmt.Errorf("cipher.NewGCM error: %w", errGCM)
	}

	return &crypt{
		aead:  aesGCM,
		block: aesBlock,
		nonce: key[len(key)-aesGCM.NonceSize():],
	}, nil
}
