package controllers

import "github.com/sniperkit/iris/mvc"

type IndexController struct{}

var indexView = mvc.View{
	Name: "index.html",
	Data: map[string]interface{}{"Title": "Home Page"},
}

func (c *IndexController) Get() mvc.View {
	return indexView
}
