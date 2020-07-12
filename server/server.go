package server

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rakyll/statik/fs"

	_ "github.com/zenzora/coveshare/statik" // Generated "Statik" fs
)

//Serve serves the server things that need serving
func Serve() {
	router := httprouter.New()

	router.POST("/api/encrypt", encrypt)
	router.GET("/d", decrypt)
	router.HandleMethodNotAllowed = false

	// This bit serves all the files in "public"
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	router.NotFound = http.FileServer(statikFS)
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), router))
}
