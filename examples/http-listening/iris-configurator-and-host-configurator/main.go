package main

import (
	"github.com/sniperkit/iris/_examples/http-listening/iris-configurator-and-host-configurator/counter"

	"github.com/sniperkit/iris"
)

func main() {
	app := iris.New()
	app.Configure(counter.Configurator)

	app.Run(iris.Addr(":8080"))
}
