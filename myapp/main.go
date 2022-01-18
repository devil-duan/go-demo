package main

import ("github.com/kataras/iris/v12"
	"myapp/router"
)

func main() {
	
	app := router.NewApp()
	app.Run(iris.Addr(":9999"))

	// m := mvc.New(demoAPI)
	// m.Handle(new(UserController))
}





