package secrets

import (
	"testing"
)


func TestGenerateNewKey(t *testing.T){
	key := GenerateNewKey()
	if len(key) != 32{
		t.Errorf("Key length should be 32 bits got %d",len(key))
	}
}
