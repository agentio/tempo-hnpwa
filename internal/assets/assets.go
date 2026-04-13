package assets

import (
	"embed"
	_ "embed"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

//go:embed css/* img/* js/*
var files embed.FS

func Handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	b, err := files.ReadFile(strings.TrimPrefix(path, "/"))
	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	if contentType := mime.TypeByExtension(filepath.Ext(path)); contentType != "" {
		w.Header().Set("content-type", contentType)
	}
	w.Write(b)
}
