package secrets

//Secrets interface defines a secret
type Secrets interface {
	Encrypt() ([]byte, error)
	Decrypt() ([]byte, error)
}

//Secret contains a secret
type Secret struct {
	PlainText  []byte
	CipherText []byte
}
