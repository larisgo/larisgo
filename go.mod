module App

go 1.12

replace (
	App/Bootstrap => ./bootstrap
	App/Config => ./config
	App/Http/Controllers => ./app/Http/Controllers
	App/Providers => ./app/Providers
	App/Routes => ./routes
)

require (
	App/Bootstrap v0.0.0-00010101000000-000000000000 // indirect
	App/Http/Controllers v0.0.0-00010101000000-000000000000 // indirect
	App/Providers v0.0.0-00010101000000-000000000000
	App/Routes v0.0.0-00010101000000-000000000000 // indirect
	github.com/larisgo/framework v0.0.0-20190610093825-683ab3c1a3cc
)
