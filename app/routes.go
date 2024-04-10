package app

import (
	"net/http"

	"github.com/EmaNymton123/PocketTroubadour/controllers"
)

func Routes() http.Handler{
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.Index)
	mux.HandleFunc("/tuto", controllers.Tuto)
	return mux
}

