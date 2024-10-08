package styleplayground

import (
	internalApp "farishadibrata.com/fiber-modular/internal/app"
	index "farishadibrata.com/fiber-modular/modules/styleplayground/routes/index"
)

func Bootstrap(app *internalApp.AppInstance) {
	group := app.Fiber.Group("/style-playground")

	group.Get("/:page", index.PlaygroundHandler)
	group.Get("/", index.PlaygroundHandler)
}
