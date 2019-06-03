package routes

import (
	. "app/controller/web"
	"github.com/larisgo/framework/Routing"
)

func Web(Route *Routing.Router) {
	Route.Get("/", IndexController.Index).Name("test")
	Route.Get("/a", IndexController.Index).Name("test")
	Route.Post("/a", IndexController.Index).Name("test")
	Route.Delete("/a", IndexController.Index).Name("test")
}
