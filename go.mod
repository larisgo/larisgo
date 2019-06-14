module App

go 1.12

replace (
	App/Http/Controllers => ./app/Http/Controllers
	Bootstrap => ./bootstrap
	Routes => ./routes
)

require (
	App/Http/Controllers v0.0.0-00010101000000-000000000000 // indirect
	Bootstrap v0.0.0-00010101000000-000000000000 // indirect
	Routes v0.0.0-00010101000000-000000000000 // indirect
	github.com/larisgo/framework v0.0.0-20190610093825-683ab3c1a3cc // indirect
)
