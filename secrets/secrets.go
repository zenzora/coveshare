package secrets

import (
	"crypto/rand"
	"errors"
	"io"
)

func PrepareSecret(msg string){
	
}

func generateUserKey() []byte{
	userKey := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, userKey); err != nil {
		panic(err.Error())
	}
	return userKey
}

func prepareMessageKey(userKey []byte, serverKey []byte) ([]byte,error){
	if len(userKey) != 32{
		return nil,errors.New("userKey must be 32 bits")
	}
	if len(serverKey) != 32{
		return nil,errors.New("serverKey must be 32 bits")
	}
	messageKey := make([]byte, 32)
	for index, element := range userKey {
		messageKey[index] = element ^ serverKey[index]
	}
	return messageKey,nil
}
