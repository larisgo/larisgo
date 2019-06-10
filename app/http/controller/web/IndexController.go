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
	return Http.NewResponse(fmt.Sprintf("%s%s", request.Get("xxx"), "123456"), 666).Header("Test", "哈哈哈")
}

func (this *Index) Show(request *Http.Request) *Http.Response {
	return Http.NewResponse(fmt.Sprintf("%s%s", request.Get("test"), "这是一个show页面"), 200).Header("Test", "哈哈哈").Header("Test", "哈哈哈1").Header("Test1", "哈哈哈")
}
