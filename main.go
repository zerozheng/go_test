package main

import (
	"test/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	mvc.Configure(app.Party("/api"), func(a *mvc.Application) {
		a.Party("/login").Handle(new(controllers.LoginController))
	})

	app.Run(iris.Addr(":8080"))
}