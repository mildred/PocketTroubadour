package controllers

import (
	"net/http"
	"fmt"
)

func Error404(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "404")
}
