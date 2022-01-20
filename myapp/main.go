package main

import ("github.com/kataras/iris/v12"
	"myapp/router"
)

func main() {
	//加载路由
	app := router.NewApp()
	app.Run(iris.Addr(":9999"))
	
}





