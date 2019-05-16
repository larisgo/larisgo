package controller

import "fmt"

type Controller struct {
}

func (this Controller) Index() {
	fmt.Println("")
	fmt.Println("Controller")
}
