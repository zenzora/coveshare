package server

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

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
	router.GET("/d/:payload", decrypt)
	router.GET("/milligram.min.css", css)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func process(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	message := []byte(r.PostFormValue("message"))
	emails := strings.Split(r.PostFormValue("allowedemails"), ",")
	expiration, err := strconv.Atoi(r.PostFormValue("expiration"))

	if err != nil {
		// Return validation error
		fmt.Fprint(w, err)
	}
	keyString := viper.GetString("key")
	keySlice, err := base64.StdEncoding.DecodeString(keyString)
	var key [32]byte
	copy(key[:], keySlice)

	if err != nil {
		fmt.Fprint(w, err)
	}
	encryptedMessage, _ := secrets.Encrypt(message, &key)
	encryptedMessageString := base64.StdEncoding.EncodeToString(encryptedMessage)

	link := "<a href=\"http://localhost:8080/d/" + url.QueryEscape(encryptedMessageString) + "\">Decrypt</a>"

	log.Println(message)
	log.Println(emails)
	log.Println(expiration)
	log.Println(encryptedMessage)

	fmt.Fprint(w, link)

}

func decrypt(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Todo error handling etc...

	payload := p.ByName("payload")
	payloadSlice, _ := base64.StdEncoding.DecodeString(payload)

	keyString := viper.GetString("key")
	keySlice, _ := base64.StdEncoding.DecodeString(keyString)
	var key [32]byte
	copy(key[:], keySlice)

	plaintext, _ := secrets.Decrypt(payloadSlice, &key)
	printme := string(plaintext)
	fmt.Fprint(w, printme)

}
