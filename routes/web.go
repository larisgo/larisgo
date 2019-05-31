package routes

import (
	. "app/controller/web"
	"github.com/larisgo/framework/Routing"
)

func Web(Route *Routing.Router) {
	Route.Get(`/aaa{x<(\w+)>}/{aaaa?}/{aaaaaaaa<\w+>?}`, IndexController{}.Index).Name("test").Domain(`a.{x}.{}.{aaaa}.a`)
	Route.Get("/a", IndexController{}.Index).Name("test")
	Route.Post("/a", IndexController{}.Index).Name("test")
	Route.Delete("/a", IndexController{}.Index).Name("test")
}
