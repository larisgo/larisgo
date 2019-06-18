package Providers

import (
	"App/Routes"
	"github.com/larisgo/framework/Contracts/Foundation"
	"github.com/larisgo/framework/Routing"
	"github.com/larisgo/framework/Support/Facades"
)

type RouteServiceProvider struct {
	App   Foundation.Application `inject:"app"`
	Defer bool                   `default:false`
}

func NewRouteServiceProvider(app Foundation.Application) (this *RouteServiceProvider) {
	this = &RouteServiceProvider{App: app}
	return this
}

/**
 * Determine if the provider is deferred.
 *
 * @return bool
 */
func (this *RouteServiceProvider) IsDeferred() bool {
	return this.Defer
}

func (this *RouteServiceProvider) Register() {
	//
}

func (this *RouteServiceProvider) Bindings() map[string]interface{} {
	return map[string]interface{}{}
}

func (this *RouteServiceProvider) Singletons() map[string]interface{} {
	return map[string]interface{}{}
}

/**
 * Bootstrap any application services.
 *
 * @return void
 */
func (this *RouteServiceProvider) Boot() {
	this.loadRoutes()

	this.App.Booted(func(app interface{}) {
		this.App.Make("router").(*Routing.Router).GetRoutes().RefreshNameLookups()
	})
}

func (this *RouteServiceProvider) loadRoutes() {
	this.mapApiRoutes()

	this.mapWebRoutes()

}

/**
 * Define the "web" routes for the application.
 *
 * These routes all receive session state, CSRF protection, etc.
 *
 * @return void
 */
func (this *RouteServiceProvider) mapWebRoutes() {
	Facades.Route().Group(map[string]string{"prefix": "/"}, Routes.Web)
}

/**
 * Define the "api" routes for the application.
 *
 * These routes are typically stateless.
 *
 * @return void
 */
func (this *RouteServiceProvider) mapApiRoutes() {
	Facades.Route().Group(map[string]string{"prefix": "api"}, Routes.Api)
}
