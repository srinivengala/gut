package api

import (
	"github.com/ant0ine/go-json-rest/rest"
)

func (api *API) Hello(w rest.ResponseWriter, r *rest.Request) {
	//log.Println("hello called")
	w.WriteJson(map[string]interface{}{
		"message": "hi there",
		"Age":     6,
	})
}
