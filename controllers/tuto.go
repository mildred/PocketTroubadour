package controllers

import (
	"net/http"
)

type TutoController struct {
	*AppContext
}

func (c *TutoController) ServeHTTP(w http.ResponseWriter, r *http.Request){
	var data struct{}
	c.Views.Lookup("tuto.html").Execute(w, data)
}

