package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/v0/", proxyHandler)
	mux.HandleFunc("/", http.FileServer(http.Dir("public")).ServeHTTP)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		p := r.URL.Path
		log.Printf("requesting %s", r.URL.Path)
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
