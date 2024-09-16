package module

import (
	"farishadibrata.com/fiber-modular/internal/app"
	"farishadibrata.com/fiber-modular/modules/styleplayground"
)

func ModuleLoader(app *app.AppInstance) {
	styleplayground.Bootstrap(app)
}
