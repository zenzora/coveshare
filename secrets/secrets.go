package secrets

import (
	"crypto/rand"
	"io"
)

func GenerateNewKey() []byte{
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err.Error())
	}
	return key
}

