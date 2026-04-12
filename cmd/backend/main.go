package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/v0/", proxyHandler)
	mux.HandleFunc("/", appHandler)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		p := r.URL.Path
		log.Printf("proxy handler %s", r.URL.Path)
		// this is actually a different API
		// base := "https://hacker-news.firebaseio.com"
		base := "https://api.hnpwa.com"
		req, err := http.NewRequest("GET", base+p, nil)
		if err != nil {
			log.Printf("%s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("%s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("response %+v", response)
		w.WriteHeader(response.StatusCode)
		_, err = io.Copy(w, response.Body)
		if err != nil {
			log.Printf("%s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "not yet", http.StatusNotFound)
}

func appHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Printf("app handler %s", r.URL.Path)
	if strings.HasPrefix(path, "/assets/") ||
		strings.HasSuffix(path, ".png") {
		http.FileServer(http.Dir("public")).ServeHTTP(w, r)
		return
	}
	b, err := os.ReadFile("public/index.html")
	if err != nil {
		log.Printf("%s", err)
	}
	w.Write(b)
}
