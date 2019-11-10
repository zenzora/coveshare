package secrets

import (
	"crypto/rand"
	"io"
)

//GenerateNewKey creates a random 32 bit key
func GenerateNewKey() []byte{
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err.Error())
	}
	return key
}

