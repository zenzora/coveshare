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

//Serve serves the server things that need serving
func Serve() {
	router := httprouter.New()
	router.GET("/", index)
	router.POST("/", process)
	// Doesn't work some times
	router.GET("/d", decrypt)
	router.GET("/milligram.min.css", css)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func process(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	if viper.GetString("encryption-type") == "aes-sha256" {
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
		secret = aess
	}

	err = secret.Encrypt()
	if err != nil {
		fmt.Fprint(w, err)
	}

	encryptedMessage := secret.GetCipherText()
	encryptedMessageString := base64.StdEncoding.EncodeToString(encryptedMessage)

	link := "<a href=\"http://localhost:8080/d?payload=" + url.QueryEscape(encryptedMessageString) + "\">Decrypt</a>"

	fmt.Fprint(w, link)

}

func decrypt(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Todo error handling etc...

	payload := r.FormValue("payload")
	payloadSlice, _ := base64.StdEncoding.DecodeString(payload)
	var secret secrets.Secrets
	// Configuration for AES-SHA256 encryption type
	if viper.GetString("encryption-type") == "aes-sha256" {
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
		secret = aess
	}

	secret.Decrypt()
	printme := string(secret.GetPlainText())
	fmt.Fprint(w, printme)

}
