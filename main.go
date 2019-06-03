package main

import (
	"bootstrap"
	"fmt"
	"github.com/larisgo/framework/Foundation/Http"
	// "github.com/larisgo/framework/Support"
	"github.com/larisgo/framework/Routing"
	"github.com/valyala/fasthttp"
	"routes"
	// "strings"
)

const LARISGO_START = 1

func main() {
	app := bootstrap.App()

	/**
	* |--------------------------------------------------------------------------
	* | Run The Application
	* |--------------------------------------------------------------------------
	* |
	* | Once we have the application, we can handle the incoming request
	* | through the kernel, and send the associated response back to
	* | the client's browser allowing them to enjoy the creative
	* | and wonderful application we have prepared for them.
	* |
	**/
	// fmt.Println(app)

	router := Routing.NewRouter()
	routes.Web(router)
	router.GetRoutes().RefreshNameLookups()
	kernel := Http.NewKernel(app, router)
	fmt.Println(`Larisgo development server started: <http://127.0.0.1:8000>`)
	// kernel.Handle() // 载入路由，配置，以及其它的服务
	panic(fasthttp.ListenAndServe(":8000", kernel.SendRequestThroughRouter))
}
