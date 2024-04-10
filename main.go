package main

import (
	"net/http"
	"log"
	"io/fs"
	"embed"

	"github.com/EmaNymton123/PocketTroubadour/controllers"

	recurparse "github.com/karelbilek/template-parse-recursive"
)

//go:embed all:public
var public_files embed.FS

//go:embed all:views
var views_files embed.FS

func run_server() error {
	// Load public assets
	public_fs, err := fs.Sub(public_files, "public")
	if err != nil {
		return err
	}

	// Load view templates
	views, err := recurparse.HTMLParseFS(nil, views_files, "*.html")
	if err != nil {
		return err
	}

	s := &http.Server{
		Addr:    ":8080",
		Handler: controllers.Routes(http.FS(public_fs), views),
	}
	log.Println("starting PocketTroubadour")
	return s.ListenAndServe()
}

func main(){
	err := run_server()
	if err != nil {
		log.Fatal(err)
	}
}
