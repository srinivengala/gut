package ui

import (
	"fmt"
	"html"
	"net/http"
	"strings"

	rice "github.com/GeertJohan/go.rice"
)

//UI is model for web gui
type UI struct {
}

//NewUI creates new web gui instance
func NewUI() *UI {
	return &UI{}
}

//Redirect to index page
func (u *UI) Redirect() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/ui/index.html", http.StatusFound) //302
	})
}

//Index page
func (u *UI) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)

		msg := fmt.Sprintf(`Hello, %q<br>
	<a href="http://localhost:9443/ui">http://localhost:9443/ui></a><br>
	<a href="http://localhost:9443/v1">http://localhost:9443/v1></a>`, html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, "<html><body><h1>%s</h1></body></html>", msg)
	})
}

//WebFolderHandler serves files in web folder
func (u *UI) WebFolderHandler() http.Handler {
	box := rice.MustFindBox("web/gut/dist/")
	return u.filterDirectoryListing(http.FileServer(box.HTTPBox()))
}

func (u *UI) filterDirectoryListing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			//fmt.Fprintf(w, "Page : %s", r.URL.Path)
			//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
