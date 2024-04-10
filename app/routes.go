package app

import (
	"net/http"

	"github.com/EmaNymton123/PocketTroubadour/controllers"
)

func Routes(assets http.FileSystem) http.Handler{
	assets_handler := http.FileServer(assets)

	mux := http.NewServeMux()
	mux.HandleFunc("/tuto", controllers.Tuto)
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path == "/" {
			controllers.Index(w, req)
			return
		}

		assets_handler.ServeHTTP(w, req)

		// Custom 404 error solution:
		// https://stackoverflow.com/questions/47285119/how-to-custom-handle-a-file-not-being-found-when-using-go-static-file-server
		// controllers.Error404(w, req)
	})
	return mux
}

