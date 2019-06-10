package routes

import (
	"github.com/larisgo/framework/Routing"
)

func Api(Route *Routing.Router) {
	Route.Get("/", nil)
	Route.Get("test", nil)
}
