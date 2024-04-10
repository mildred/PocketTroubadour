package controllers

import (
	"net/http"
	"html/template"
)

func Index(w http.ResponseWriter, r *http.Request){
	var data struct{}
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w,data)
}
