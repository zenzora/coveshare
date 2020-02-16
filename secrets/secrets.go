package secrets

type Secrets interface {
	Encrypt() []byte, error
	Decrypt() []byte, error
}

type Secret struct {
	plainText []byte
	cipherText []byte
}
