package server


import (
    "fmt"
    "net/http"
    "log"

    "github.com/julienschmidt/httprouter"
)

//Serve serves the server things that need serving
func Serve() {
    router := httprouter.New()
	router.GET("/", index)
	router.POST("/",process)
    router.GET("/milligram.min.css", css)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func process(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	fmt.Fprint(w,r.PostFormValue("message"))
}

