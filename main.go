package main

import (
	"net/http"
	"log"
	"io/fs"
	"embed"
	"flag"

	"github.com/EmaNymton123/PocketTroubadour/controllers"

	recurparse "github.com/karelbilek/template-parse-recursive"
)

//go:embed all:public
var public_files embed.FS

//go:embed all:views
var views_files embed.FS

func run_server() error {
	var err error
	var public_dir, views_dir, listen string
	flag.StringVar(&listen, "listen", ":8080", "Address to listen to")
	flag.StringVar(&public_dir, "public", "", "Location for public dir")
	flag.StringVar(&views_dir, "views", "", "Location for views dir")
	flag.Parse()

	app := new(controllers.AppContext)

	// Load public assets
	if public_dir == "" {
		public_fs, err := fs.Sub(public_files, "public")
		if err != nil {
			return err
		}
		app.PublicFS = http.FS(public_fs)
	} else {
		app.PublicFS = http.Dir(public_dir)
	}

	// Load view templates
	if views_dir == "" {
		views_fs, err := fs.Sub(views_files, "views")
		if err != nil {
			return err
		}
		app.Views, err = recurparse.HTMLParseFS(nil, views_fs, "*.html")
		if err != nil {
			return err
		}
	} else {
		app.Views, err = recurparse.HTMLParse(nil, views_dir, "*.html")
		if err != nil {
			return err
		}
	}

	s := &http.Server{
		Addr:    listen,
		Handler: app.Routes(),
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
