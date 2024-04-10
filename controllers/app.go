package controllers

import (
	"net/http"
	"html/template"
)

type AppContext struct {
	Views *template.Template
}

func Routes(assets http.FileSystem, views *template.Template) http.Handler {
	assets_handler := http.FileServer(assets)

	app := &AppContext{
		Views: views,
	}

	mux := http.NewServeMux()
	mux.Handle("/tuto", &TutoController{app})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path == "/" {
			index := &IndexController{app}
			index.ServeHTTP(w, req)
			return
		}

		assets_handler.ServeHTTP(w, req)

		// Custom 404 error solution:
		// https://stackoverflow.com/questions/47285119/how-to-custom-handle-a-file-not-being-found-when-using-go-static-file-server
		// Error404(w, req)
	})
	return mux
}



