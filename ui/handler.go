package ui

import (
	"fmt"
	"html"
	"net/http"
)

func (u *UI) changeMe() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)

		msg := fmt.Sprintf(`Hello, %q<br>
	<a href="http://localhost:9443/ui">Goto UI</a><br>
	<a href="http://localhost:9443/v1">Goto API</a>`, html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, "<html><body><h1>%s</h1></body></html>", msg)
	})
}
