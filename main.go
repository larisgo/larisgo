package main

import (
	. "app/controller/web"
	"bootstrap"
	"fmt"
	// "github.com/larisgo/larisgo/Foundation/Http"
	// "net/http"
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
	// kernel := Http.NewKernel(app)
	// // kernel.handle(router)
	// // kernel.Terminate(nil, nil)
	fmt.Printf("%+v", app)
	// fmt.Printf("%+v", kernel)
	action := new(IndexController).Index
	fmt.Print(action)
	action()
	new(IndexController).A()
	// http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	// })

	// panic(http.ListenAndServe(":8000", nil))
}
