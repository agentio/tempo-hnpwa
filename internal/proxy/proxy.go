package proxy

import (
	"io"
	"net/http"
)

const APIServer = "https://api.hnpwa.com"

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		req, err := http.NewRequest("GET", APIServer+r.URL.Path, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(response.StatusCode)
		_, err = io.Copy(w, response.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
}
