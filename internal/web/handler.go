package web

import (
	"encoding/json"
	"log"
	"net/http"
	"snip/internal/service"
	"snip/internal/storage"
	"strings"
)

var body struct {
	URL string `json:"url"`
}

type res struct {
	Code string `json:"code"`
	URL  string `json:"url"`
}

func Redirect(w http.ResponseWriter, r *http.Request) {

	code := strings.TrimPrefix(r.URL.Path, "/")
	if code == "favicon.ico" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	url, err := storage.Get(code)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	log.Printf("redirect: %s -> %s (ip: %s)", code, url.URL, r.RemoteAddr)

	errClick := storage.AddClick(code)
	if errClick != nil {
		log.Printf("failed to add click for code %s: %v", code, errClick)
	}

	http.Redirect(w, r, url.URL, http.StatusFound)

}

func CreateLinkPOST(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	code := service.GenCode()

	errSave := storage.Save(code, body.URL)
	if errSave != nil {
		log.Printf("failed to save URL %s: %v", body.URL, errSave)
	}

	log.Printf("new link created: %s -> %s", code, body.URL)

	link := "http://localhost:6767/" + code

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res{Code: code, URL: link})

}

func GetLinkGET(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	code := strings.TrimPrefix(r.URL.Path, "/api/links/")
	if code == "favicon.ico" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	url, err := storage.Get(code)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	log.Printf("link info requested: %s", code)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(url)

}
