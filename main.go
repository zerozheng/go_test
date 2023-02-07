package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	mvc.Configure(app.Party("/api"), func(a *mvc.Application) {
		a.Party("/login").Handle(new(LoginController))
	})

	app.Run(iris.Addr(":8080"))
}

type LoginController struct {
	Ctx iris.Context
}

func (c *LoginController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/signout", "LoginHandle")

	println("I am a BeforeActivation")
}

func (l *LoginController) LoginHandle() string {
	println("I am a LoginHandle")

	return "LoginHandle success"
}

func (l *LoginController) Post() string {
	println("I am a post")

	return "post success"
}
