package main

import (
	"github.com/MarErm27/GoAdmin/context"
	"github.com/MarErm27/GoAdmin/modules/auth"
	c "github.com/MarErm27/GoAdmin/modules/config"
	"github.com/MarErm27/GoAdmin/modules/db"
	"github.com/MarErm27/GoAdmin/modules/service"
	"github.com/MarErm27/GoAdmin/plugins"
)

type Example struct {
	*plugins.Base
}

var Plugin = &Example{
	Base: &plugins.Base{PlugName: "example"},
}

func (example *Example) InitPlugin(srv service.List) {
	example.InitBase(srv, "example")
	Plugin.App = example.initRouter(c.Prefix(), srv)
}

func (example *Example) initRouter(prefix string, srv service.List) *context.App {

	app := context.NewApp()
	route := app.Group(prefix)
	route.GET("/example", auth.Middleware(db.GetConnection(srv)), example.TestHandler)

	return app
}

func (example *Example) TestHandler(ctx *context.Context) {

}
