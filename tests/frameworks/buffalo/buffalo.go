package buffalo

import (
	// add buffalo adapter
	_ "github.com/MarErm27/GoAdmin/adapter/buffalo"
	"github.com/MarErm27/GoAdmin/modules/config"
	"github.com/MarErm27/GoAdmin/modules/language"
	"github.com/MarErm27/GoAdmin/plugins/admin/modules/table"
	"github.com/MarErm27/themes/adminlte"

	// add mysql driver
	_ "github.com/MarErm27/GoAdmin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "github.com/MarErm27/GoAdmin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "github.com/MarErm27/GoAdmin/modules/db/drivers/sqlite"
	// add mssql driver
	_ "github.com/MarErm27/GoAdmin/modules/db/drivers/mssql"
	// add adminlte ui theme
	_ "github.com/MarErm27/themes/adminlte"

	"github.com/MarErm27/GoAdmin/template"
	"github.com/MarErm27/GoAdmin/template/chartjs"

	"net/http"
	"os"

	"github.com/MarErm27/GoAdmin/engine"
	"github.com/MarErm27/GoAdmin/plugins/admin"
	"github.com/MarErm27/GoAdmin/plugins/example"
	"github.com/MarErm27/GoAdmin/tests/tables"
	"github.com/gobuffalo/buffalo"
)

func newHandler() http.Handler {
	bu := buffalo.New(buffalo.Options{
		Env:  "test",
		Addr: "127.0.0.1:9033",
	})

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)

	examplePlugin := example.NewExample()

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(bu); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	bu.ServeFiles("/uploads", http.Dir("./uploads"))

	return bu
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	bu := buffalo.New(buffalo.Options{
		Env:  "test",
		Addr: "127.0.0.1:9033",
	})

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(gens)

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfig(config.Config{
		Databases: dbs,
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:    language.EN,
		IndexUrl:    "/",
		Debug:       true,
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}).
		AddPlugins(adminPlugin).Use(bu); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	bu.ServeFiles("/uploads", http.Dir("./uploads"))

	return bu
}
