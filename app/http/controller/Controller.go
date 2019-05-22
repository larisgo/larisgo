package controller

import "fmt"

type Controller struct {
}

func (Controller) Index() {
	fmt.Println("")
	fmt.Println("Controller")
}
