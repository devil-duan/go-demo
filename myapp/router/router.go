package router

import (
	"myapp/controller"
	"net/http"
	"strings"

	"github.com/kataras/iris/v12"
)


func NewApp() *iris.Application {
    app := iris.New()

    app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
        ctx.HTML("<b>Resource Not found</b>")
    })

    app.Get("/profile/{username}", func(ctx iris.Context) {
        ctx.Writef("Hello %s", ctx.Params().Get("username"))
    })
	
	app.Post("/demo",controller.Create)
	app.Get("/demo", controller.List)
    app.HandleDir("/", "./public")

    myOtherHandler := func(ctx iris.Context) {
        ctx.Writef("inside a handler which is fired manually by our custom router wrapper")
    }

    // 用本地 net/http 处理程序包装路由器。
    // 如果 url 不包含任何 "." (即: .css, .js...)
    // (取决于应用程序，你可能需要添加更多文件服务器异常,
    // 然后处理程序将执行负责
    // 已注册的路由 (看上去是 "/" 和 "/profile/{username}")
    // 如果没有， 则它将基于根 "/" 路径文件提供服务。
    app.WrapRouter(func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
        path := r.URL.Path

        if strings.HasPrefix(path, "/other") {
                // 获取并释放上下文以便使用它来执行
                // 我们的自定义处理程序
                // 记住：我们使用 net/http.Handler 是因为
                // 我们在路由器本身之前处于“低级”状态。
                ctx := app.ContextPool.Acquire(w, r)
                myOtherHandler(ctx)
                app.ContextPool.Release(ctx)
                return
            }

            // 否则继续照常服务路由。
            router.ServeHTTP(w, r) 
    })

    return app
}



