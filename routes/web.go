package routes

import (
	. "app/controller/web"
	"github.com/larisgo/framework/Routing"
)

func Web(Route *Routing.Router) {
	Route.Group(map[string]string{"prefix": "hah"}, func(Route *Routing.Router) {
		Route.Get("/", IndexController.Show)
	})
	Route.Get("a/{xxx}/{ssss?}", IndexController.Index)
	Route.Get("json", IndexController.Json)
	Route.Post("a", IndexController.Index)
	Route.Delete("a", IndexController.Index)
}
