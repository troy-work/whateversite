package main

import (
	"net/http"
	. "github.com/troy-work/whateversite/components"
	"log"
)

func redirHttps(w http.ResponseWriter, req *http.Request) {
	log.Printf("redirecting to https")
	http.Redirect(w, req, "https://localhost:8081"+req.RequestURI, http.StatusMovedPermanently)
}


func main() {

	http.HandleFunc("/", DefaultHandler)
	http.HandleFunc("/view/", MakeHandler(ViewHandler))
	http.HandleFunc("/edit/", MakeHandler(EditHandler))
	http.HandleFunc("/save/", MakeHandler(SaveHandler))

	err := http.ListenAndServeTLS(":8081", "cert/server.pem", "cert/server.key", nil)
	if (err != nil){
		log.Fatal("ListenAndServe", err)
	}
	http.ListenAndServe(":8080", http.HandlerFunc(redirHttps))

}