package main

import (
	"net/http"
	"log"
	"github.com/EmaNymton123/PocketTroubadour/controllers"
)

func main(){
	s := &http.Server{
		Addr:           ":8080",
		Handler:        controllers.Routes(),
	}
	log.Println("starting PocketTroubadour")
	log.Fatal(s.ListenAndServe())
}
