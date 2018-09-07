package ui

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/srinivengala/gut/api"
)

// StartServer starts gut server
func StartServer() int {
	restAPI := rest.NewApi()
	//TODO api.Use(rest.DefaultProdStack...)
	restAPI.Use(rest.DefaultDevStack...)

	a := api.NewAPI()
	restRouter, _ := a.MakeRouter()

	restAPI.SetApp(restRouter)

	http.Handle("/v1/", http.StripPrefix("/v1", restAPI.MakeHandler()))

	webUI := NewUI()

	//NOTE: ending with slash is important
	http.Handle("/ui/", http.StripPrefix("/ui/", webUI.WebFolderHandler()))
	http.Handle("/", webUI.Redirect())

	//TODO http.ListenAndServeTLS
	if err := http.ListenAndServe(":9443", nil); err != nil {
		return 1
	}
	return 0
}
