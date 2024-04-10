package controllers

import (
	"net/http"
	"html/template"
)

func Tuto(w http.ResponseWriter, r *http.Request){
	var data struct{}
	tmpl := template.Must(template.ParseFiles("tuto.html"))
	tmpl.Execute(w,data)
}

