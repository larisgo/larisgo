package main

import (
	"Bootstrap"
	"github.com/larisgo/framework/Contracts/Http"
)

const LARISGO_START = 1

func main() {
	app := Bootstrap.App()

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

	kernel := app.Make("kernel").(Http.Kernel)

	kernel.Handle()
}
