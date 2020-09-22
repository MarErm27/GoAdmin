package example

import (
	c "github.com/MarErm27/GoAdmin/modules/config"
	"github.com/MarErm27/GoAdmin/modules/service"
	"github.com/MarErm27/GoAdmin/plugins"
)

type Example struct {
	*plugins.Base
}

func NewExample() *Example {
	return &Example{
		Base: &plugins.Base{PlugName: "example"},
	}
}

func (e *Example) InitPlugin(srv service.List) {
	e.InitBase(srv, "example")
	e.App = e.initRouter(c.Prefix(), srv)
}
