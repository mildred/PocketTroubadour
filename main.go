package main

import (
	"net/http"
	"log"
	"io/fs"
	"embed"

	"github.com/EmaNymton123/PocketTroubadour/app"
)

//go:embed all:public
var public_files embed.FS

func PublicFS() http.FileSystem {
	res, err := fs.Sub(public_files, "public")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(res)
}

func main(){
	s := &http.Server{
		Addr:    ":8080",
		Handler: app.Routes(PublicFS()),
	}
	log.Println("starting PocketTroubadour")
	log.Fatal(s.ListenAndServe())
}
