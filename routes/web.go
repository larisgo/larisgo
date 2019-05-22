package routes

import (
	. "app/controller/web"
	"github.com/larisgo/larisgo/Routing"
)

func Web(Route *Routing.Router) {
	Route.Get(`/aaa/{aaaa?}/{aaaaaaaa<\w+>?}`, IndexController{}.Index).Name("test").Matches(true, true)
	Route.Get("/a", IndexController{}.Index).Name("test")
	Route.Post("/a", IndexController{}.Index).Name("test")
	Route.Delete("/a", IndexController{}.Index).Name("test")
}
