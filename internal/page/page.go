package page

import (
	"net/http"

	. "github.com/delaneyj/gostar/elements"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if err := index().Render(w); err != nil {
		return
	}
}

func index() ElementRenderer {
	return Group(
		Text("<!DOCTYPE html>"),
		HTML(
			HEAD(
				LINK().REL("icon").TYPE("image/png").HREF("/img/icon-512x512.png"),
				LINK().REL("stylesheet").HREF("/css/index.css"),
				META().CHARSET("UTF-8"),
				META().NAME("description").CONTENT("Tempo implementation of a Hacker News browser."),
				META().NAME("viewport").CONTENT("width=device-width, initial-scale=1.0"),
				SCRIPT().TYPE("module").SRC("/js/index.js"),
				TITLE().Text("Tempo + HNPWA"),
			),
			BODY(
				DIV().ID("app"),
			),
		).LANG("en"),
	)
}
