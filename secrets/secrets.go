package secrets

//Secrets interface can encrypt and decrypt
type Secrets interface {
	Encrypt() ([]byte, error)
	Decrypt() ([]byte, error)
}

//Secret contains a secret
type Secret struct {
	PlainText  []byte
	CipherText []byte
}
