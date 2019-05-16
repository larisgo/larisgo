package web

import (
	"app/controller"
	"fmt"
)

type IndexController struct {
	controller.Controller
}

func (this IndexController) A() {
	fmt.Println("")
	fmt.Println("A")
}
