package main

import (
	"myapp/model"

	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	demoAPI := app.Party("/demo")
	{
		demoAPI.Use(iris.Compression)

		demoAPI.Get("/", list)

		demoAPI.Post("/", create)
	}
	app.Listen(":9999")

	m := mvc.New(demoAPI)
	m.Handle(new(UserController))
}

func list(ctx iris.Context) {

	users := []model.User{
		{"张三"},
		{"李四"},
		{"王五"},
	}
	ctx.JSON(users)
}

func create(ctx iris.Context) {
	var user model.User
	err := ctx.ReadJSON(&user)

	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title("create user fail").DetailErr(err))
		return
	}
	println("Received User: " + user.Name)

	ctx.StatusCode(iris.StatusCreated)

}

type UserController struct {
	/* dependencies */
}

func (c *UserController) Get() []model.User {
	return []model.User{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}
}

func (c *UserController) Post(b model.User) int {
	println("Received Book: " + b.Name)

	return iris.StatusCreated
}
