package controllers

import (
	"github.com/sniperkit/iris/_examples/http_responsewriter/quicktemplate/templates"

	"github.com/sniperkit/iris"
)

// ExecuteTemplate renders a "tmpl" partial template to the `context#ResponseWriter`.
func ExecuteTemplate(ctx iris.Context, tmpl templates.Partial) {
	ctx.Gzip(true)
	ctx.ContentType("text/html")
	templates.WriteTemplate(ctx.ResponseWriter(), tmpl)
}
