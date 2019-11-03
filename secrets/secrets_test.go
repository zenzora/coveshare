package secrets

import (
	"reflect"
	"testing"
)


func TestGenerateUserKey(t *testing.T){
	key := generateUserKey()
	if len(key) != 32{
		t.Errorf("Key length should be 32 bits got %d",len(key))
	}
}

func TestPrepareMessageKey(t *testing.T){
	validKey := []byte{11, 101, 236, 65, 177, 1, 124, 58, 9, 40, 85, 245, 143, 179, 115, 193, 177, 28, 161, 207, 75, 243, 235, 177, 0, 230, 243, 94, 255, 60, 88, 197}
	shortKey := []byte{11, 236, 65, 177, 1,124, 58, 9, 40, 85, 245, 143, 179, 115, 193, 177, 28, 161, 207, 75, 243, 235, 177, 0, 230, 243, 94, 255, 60, 88, 197}

	//Test identical keys
	messageKey,err := prepareMessageKey(validKey, validKey)
	if err != nil{
		t.Errorf("Testing identical keys returned error: %s",err)
	}
	if !reflect.DeepEqual(messageKey,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}) {
		t.Errorf("Testing identical keys did not return all zeros")
	}

	//Test with short keys
	_,err = prepareMessageKey(shortKey, shortKey)
	if err == nil{
		t.Errorf("Testing short keys did not result in an error")
	}

}