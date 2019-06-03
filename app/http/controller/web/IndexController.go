package web

import (
	. "app/controller"
	"fmt"
	"github.com/larisgo/framework/Http"
)

type Index struct {
	*Controller
}

var IndexController *Index = &Index{}

func (this *Index) Index(*Http.Request) *Http.Response {
	fmt.Println("")
	fmt.Println("A")
	return Http.NewResponse("123456", 200)
}
