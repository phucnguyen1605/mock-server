package main

import (
	"github.com/phucnh/api-mock/src/app"
	"github.com/phucnh/api-mock/src/handlers"
	"github.com/phucnh/api-mock/src/middleware"
)

func main() {
	app := app.New()
	setEnvs(app)
	startHTTPServer(app)
}

func setEnvs(app *app.App) {
	app.SetEnv("SERVER_PORT", "5151")
	app.SetEnv("JWT_KEY", "wqGyEBBfPK9w3Lxw")
	app.SetEnv("DB_CONNECTION_STRING", `mock_api:mock_api@/mock_api`)
}

func startHTTPServer(app *app.App) {
	handlers.AddRouters(app)
	app.UseAppMiddlewareFunc(middleware.AuthMiddleWareFunc())
	app.Handle()
}
