package secrets

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

//GenerateNewKey creates a random 32 bit key
func GenerateNewKey() *[32]byte {
	key := [32]byte{}
	if _, err := io.ReadFull(rand.Reader, key[:]); err != nil {
		panic(err.Error())
	}
	return &key
}

//AesSecret is a type of secret, unlike KMS it needs a key
type AesSecret struct {
	Secret
	Key *[32]byte
}

// Encrypt function copied from: https://github.com/gtank/cryptopasta - CC License
func (aess AesSecret) Encrypt() (ciphertext []byte, err error) {
	block, err := aes.NewCipher(aess.Key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, aess.PlainText, nil), nil
}

// Decrypt function copied from: https://github.com/gtank/cryptopasta - CC License
func (aess AesSecret) Decrypt() (plaintext []byte, err error) {
	block, err := aes.NewCipher(aess.Key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(aess.CipherText) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(nil,
		aess.CipherText[:gcm.NonceSize()],
		aess.CipherText[gcm.NonceSize():],
		nil,
	)
}
