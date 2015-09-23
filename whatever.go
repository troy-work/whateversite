package main

import (
	"net/http"
	. "github.com/troy-work/whateversite/components"
	. "gopkg.in/tylerb/graceful.v1"
	"log"
	"time"
)

func redirectHttps(w http.ResponseWriter, req *http.Request) {
	log.Printf("redirecting to https")
	http.Redirect(w, req, "https://localhost:1443"+req.RequestURI, http.StatusMovedPermanently)
}


func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", DefaultHandler)
	mux.HandleFunc("/view/", MakeHandler(ViewHandler))
	mux.HandleFunc("/edit/", MakeHandler(EditHandler))
	mux.HandleFunc("/save/", MakeHandler(SaveHandler))

	srvTls :=  &http.Server{
		Addr: ":1443",
		Handler: mux,
	}

	err := ListenAndServeTLS(srvTls, "cert/server.pem", "cert/server.key", time.Second)
	if (err != nil){
		log.Fatal("ListenAndServe", err)
	}

	mux2 := http.NewServeMux()
	mux2.Handle("*",http.HandlerFunc(redirectHttps))

	srv :=  &http.Server{
		Addr: ":1443",
		Handler: mux2,
	}

	//DON'T USE ListenAndServe like this. You can't shut it down
	ListenAndServe(srv, time.Second)

}