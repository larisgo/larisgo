package main

import (
	"bootstrap"
	"fmt"
	"github.com/larisgo/framework/Foundation/Http"
	// "github.com/larisgo/framework/Support"
	"github.com/larisgo/framework/Routing"
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
	router.Group(map[string]string{"prefix": "api"}, func(router *Routing.Router) {
		routes.Api(router)
	})
	routes.Web(router)
	router.GetRoutes().RefreshNameLookups()
	kernel := Http.NewKernel(app, router)
	fmt.Println(`Larisgo development server started: <http://127.0.0.1:8000>`)
	kernel.Handle()
}
