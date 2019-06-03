package controller

import (
	"fmt"
	"github.com/larisgo/framework/Http"
)

type Controller struct {
}

func (*Controller) Index(*Http.Request) *Http.Response {
	fmt.Println("")
	fmt.Println("Controller")
	return nil
}
