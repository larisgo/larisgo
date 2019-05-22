package web

import (
	. "app/controller"
	"fmt"
)

type IndexController struct {
	Controller
}

func (IndexController) Index() {
	fmt.Println("")
	fmt.Println("A")
}
