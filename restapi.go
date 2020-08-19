package restapi

import (
	"github.com/vulcangz/restapi/middleware"
	"github.com/vulcangz/restapi/system"
)

func NewApplication(configFile string) *system.Application {
	app := new(system.Application)
	app.Init(configFile)

	// add default middleware
	app.Container.Filter(middleware.RequestID)
	app.Container.Filter(middleware.Logger)
	app.Container.Filter(middleware.Recoverer)
	app.Container.Filter(app.Plugins)

	return app
}
