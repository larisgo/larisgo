module app

go 1.12

replace (
	app/controller => ./app/http/controller
	bootstrap => ./bootstrap
	routes => ./routes
)

require (
	app/controller v0.0.0-00010101000000-000000000000 // indirect
	bootstrap v0.0.0-00010101000000-000000000000
	github.com/larisgo/framework v0.0.0-20190531075953-06369ea13a66
	github.com/valyala/fasthttp v1.3.0 // indirect
	routes v0.0.0-00010101000000-000000000000
)
