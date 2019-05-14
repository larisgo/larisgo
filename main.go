package main

import (
	// "fmt"
	"framework/bootstrap"
	"github.com/larisgo/larisgo/Http"
	"time"
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
	kernel := app.Make()
	response := kernel.Handle(*Http.Request.capture())
	response.Send()
	kernel.Terminate(response, request)
	// fmt.Println(app)
}
