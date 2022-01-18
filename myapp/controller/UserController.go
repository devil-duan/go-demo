package controller

import (
	"myapp/model"

	"github.com/kataras/iris/v12"
)

 func List(ctx iris.Context) {

	users := []model.User{
		{"张三"},
		{"李四"},
		{"王五"},
	}
	ctx.JSON(users)
}

func Create(ctx iris.Context) {
	var user model.User
	err := ctx.ReadJSON(&user)

	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title("create user fail").DetailErr(err))
		return
	}
	println("Received User: " + user.Name)

	ctx.StatusCode(iris.StatusCreated)

}

