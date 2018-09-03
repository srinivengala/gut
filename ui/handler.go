package ui

import (
	"fmt"
	"html"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	msg := fmt.Sprintf(`Hello, %q<br>
	<a href="http://localhost:9443/ui">http://localhost:9443/ui></a><br>
	<a href="http://localhost:9443/v1">http://localhost:9443/v1></a>`, html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "<html><body><h1>%s</h1></body></html>", msg)
}
