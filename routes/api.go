package Routes

import (
	. "App/Http/Controllers/Web"
	"github.com/larisgo/framework/Routing"
)

func Api(Route *Routing.Router) {
	Route.Get("json", IndexController.Json)
	Route.Get("test", nil)
}
