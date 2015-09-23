package main

import (
	"net/http"
	. "github.com/troy-work/whateversite/components"
	. "gopkg.in/tylerb/graceful.v1"
	"log"
)

func redirectHttps(w http.ResponseWriter, req *http.Request) {
	log.Printf("redirecting to https")
	http.Redirect(w, req, "https://localhost:1443"+req.RequestURI, http.StatusMovedPermanently)
}


func main() {

	go func() {
		mux := http.NewServeMux()

		mux.HandleFunc("/", DefaultHandler)
		mux.HandleFunc("/view/", MakeHandler(ViewHandler))
		mux.HandleFunc("/edit/", MakeHandler(EditHandler))
		mux.HandleFunc("/save/", MakeHandler(SaveHandler))

		srvTls := &Server{
			Timeout: 1,
			Server:  &http.Server{Addr: ":1443", Handler: mux},
		}
		log.Printf("Starting https listener")
		err := srvTls.ListenAndServeTLS("cert/server.pem", "cert/server.key")

		if (err != nil) {
			log.Fatal("Tls ", err)
		}
	}()

	srv := &Server{
		Timeout: 1,
		Server:  &http.Server{Addr: ":8080", Handler: http.HandlerFunc(redirectHttps)},
	}
	log.Printf("Starting http server")
	err := srv.ListenAndServe()
	if (err != nil){
		log.Fatal("http ", err)
	}

}