package controllers

import (
	"net/http"
)

type IndexController struct {
	*AppContext
}

func (c *IndexController) ServeHTTP(w http.ResponseWriter, r *http.Request){
	var data struct{}
	c.Views.Lookup("views/index.html").Execute(w, data)
}
