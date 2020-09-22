package example

import (
	"github.com/MarErm27GoAdmin/context"
	"github.com/MarErm27GoAdmin/modules/auth"
	"github.com/MarErm27GoAdmin/modules/db"
	"github.com/MarErm27GoAdmin/modules/service"
)

func (e *Example) initRouter(prefix string, srv service.List) *context.App {

	app := context.NewApp()
	route := app.Group(prefix)
	route.GET("/example", auth.Middleware(db.GetConnection(srv)), e.TestHandler)

	return app
}
