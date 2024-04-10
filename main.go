package main

import (
	"net/http"
	"log"

	"github.com/EmaNymton123/PocketTroubadour/app"
)

func main(){
	s := &http.Server{
		Addr:    ":8080",
		Handler: app.Routes(),
	}
	log.Println("starting PocketTroubadour")
	log.Fatal(s.ListenAndServe())
}
