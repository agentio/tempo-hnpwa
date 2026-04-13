package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/agentio/tempo-hnpwa/internal/assets"
	"github.com/agentio/tempo-hnpwa/internal/page"
	"github.com/agentio/tempo-hnpwa/internal/proxy"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("serving on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, &App{}))
}

type App struct{}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Printf("%s", path)
	switch {
	case strings.HasPrefix(path, "/v0/"):
		proxy.Handler(w, r)
	case strings.HasPrefix(path, "/img/"), strings.HasPrefix(path, "/js/"), strings.HasPrefix(path, "/css/"):
		assets.Handler(w, r)
	default:
		page.Handler(w, r)
	}
}
