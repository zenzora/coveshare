package server

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"github.com/zenzora/coveshare/secrets"
)

func encrypt(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Verify parameters
	message := []byte(r.PostFormValue("message"))
	//emails := strings.Split(r.PostFormValue("allowedemails"), ",")
	_, err := strconv.Atoi(r.PostFormValue("expiration"))
	if err != nil {
		// Return validation error
		fmt.Fprint(w, err)
	}

	var secret secrets.Secrets

	// Configuration for AES-SHA256 encryption type
	if viper.GetString("encryption_type") == "aes-sha256" {
		var aess secrets.AesSha256Secret
		aess.PlainText = message

		keyString := viper.GetString("key")
		keySlice, err := base64.StdEncoding.DecodeString(keyString)
		var key [32]byte
		copy(key[:], keySlice)
		aess.Key = &key
		if err != nil {
			fmt.Fprint(w, err)
		}
		secret = &aess
	}
	err = secret.Encrypt()

	if err != nil {
		fmt.Fprint(w, err)
	}

	encryptedMessage := secret.GetCipherText()
	encryptedMessageString := base64.StdEncoding.EncodeToString(encryptedMessage)

	link := viper.GetString("base_url") + "/d?payload=" + url.QueryEscape(encryptedMessageString)
	fmt.Fprint(w, link)
}

func decrypt(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Todo error handling etc...

	payload := r.FormValue("payload")
	payloadSlice, _ := base64.StdEncoding.DecodeString(payload)
	var secret secrets.Secrets
	// Configuration for AES-SHA256 encryption type
	if viper.GetString("encryption_type") == "aes-sha256" {
		var aess secrets.AesSha256Secret
		aess.CipherText = payloadSlice

		keyString := viper.GetString("key")
		keySlice, err := base64.StdEncoding.DecodeString(keyString)
		var key [32]byte
		copy(key[:], keySlice)
		aess.Key = &key
		if err != nil {
			fmt.Fprint(w, err)
		}
		secret = &aess
	}

	err := secret.Decrypt()
	if err != nil {
		log.Println(err)
	}
	printme := string(secret.GetPlainText())
	fmt.Fprint(w, printme)

}
