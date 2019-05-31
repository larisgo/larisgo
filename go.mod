module app

go 1.12

replace (
	app/controller => ./app/http/controller
	bootstrap => ./bootstrap
	routes => ./routes
)
