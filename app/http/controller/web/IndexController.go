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

func (this *Index) Index(request *Http.Request) *Http.Response {
	return Http.NewResponse(fmt.Sprintf("%s%s", request.Get("xxx"), "123456"), 200)
}
