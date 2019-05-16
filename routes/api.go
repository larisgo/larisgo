package routes

import (
	"github.com/larisgo/larisgo/Routing"
)

func Web(Route *Routing.Router) {
	Route.Get("/", nil)
	Route.Group("/aaa", func(Route *Routing.Router) {
		Route.Get("/a", nil)
	})
	Route.Get("/a", nil)
}
