package controllers

import (
	"net/http"
	"html/template"
)

func index(w http.ResponseWriter, r *http.Request){
	var data struct{}
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w,data)
}

func tuto(w http.ResponseWriter, r *http.Request){
	var data struct{}
	tmpl := template.Must(template.ParseFiles("tuto.html"))
	tmpl.Execute(w,data)
}

func Routes() http.Handler{
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/tuto", tuto)
	return mux
}
