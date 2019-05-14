package main

import (
	"fmt"
	"main/bootstrap"
)

func main() {
	app := bootstrap.App()
	fmt.Printf("%+v", app)
	fmt.Println("Hello world")
}
