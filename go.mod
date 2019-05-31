module app

go 1.12

require (
	app/controller v0.0.0-00010101000000-000000000000 // indirect
	bootstrap v0.0.0-00010101000000-000000000000 // indirect
	github.com/larisgo/larisgo v0.0.0-20190515024325-9ac19b75ed26
	routes v0.0.0-00010101000000-000000000000 // indirect
)

replace (
	app/controller => ./app/http/controller
	bootstrap => ./bootstrap
	routes => ./routes
)
