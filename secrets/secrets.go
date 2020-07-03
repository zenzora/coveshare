package secrets

import "time"

//Secrets interface can encrypt and decrypt
type Secrets interface {
	Encrypt() error
	Decrypt() error
	GetCipherText() []byte
	GetPlainText() []byte
	//SetPlainText([]byte) error
	//SetExpiration(time.Time) error
	//Sign() ([]byte, error)
	//Verify() (bool, error)
}

//Secret contains a secret
type Secret struct {
	PlainText  []byte
	CipherText []byte
	Expiration time.Time
	//Emails     []string
	Signature []byte
}
