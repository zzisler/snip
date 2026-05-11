package web

import (
	"log"
	"net/http"
	"os"
)

func Start() {

	port := os.Getenv("PORT")

	log.Printf("starting server on :%s", port)

	http.HandleFunc("/", Redirect)

	http.HandleFunc("/api/links", CreateLinkPOST)

	http.HandleFunc("/api/links/{code}", GetLinkGET)

	http.ListenAndServe(":"+port, nil)
}
