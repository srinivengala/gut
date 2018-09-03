package gut

import (
	"bytes"
	"compress/gzip"
	"net/http"

	"github.com/srinivengala/gut/ui"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/srinivengala/gut/api"
)

// StartServer starts gut server
func StartServer() int {
	restAPI := rest.NewApi()
	//TODO api.Use(rest.DefaultProdStack...)
	restAPI.Use(rest.DefaultDevStack...)

	// restRouter, err := rest.MakeRouter(
	// 	rest.Get("/v1/hello", helloRestHandler)
	// )

	restRouter, _ := api.NewRouter()

	restAPI.SetApp(restRouter)

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	//mux := http.NewServeMux()
	//NOTE: ending with slash is important
	http.Handle("/v1/", http.StripPrefix("/v1", restAPI.MakeHandler()))
	http.Handle("/ui/", http.StripPrefix("/ui", http.HandlerFunc(ui.Handler)))
	http.HandleFunc("/", ui.Handler)

	if err := http.ListenAndServe(":9443", nil); err != nil {
		return 1
	}
	return 0
}

// func helloRestHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.Header().Set("Content-Encoding", "gzip")

// 	writer, err := gzip.NewWriterLevel(w, gzip.BestCompression)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	defer writer.Close()

// 	body := []byte("{\"message\":\"Hi there\"}")
// 	if _, err := writer.Write(body); err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	//fmt.Fprintf(w, &[]byte("{\"message\":\"Hi there\"}"))
// }

func doGzip(a *[]byte) ([]byte, error) {
	var b bytes.Buffer
	gz, _ := gzip.NewWriterLevel(&b, gzip.BestSpeed)
	defer gz.Close()
	if _, err := gz.Write(*a); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}