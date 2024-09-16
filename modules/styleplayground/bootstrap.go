package styleplayground

import (
	internalApp "farishadibrata.com/fiber-modular/internal/app"
	view "farishadibrata.com/fiber-modular/modules/styleplayground/view"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func Bootstrap(app *internalApp.AppInstance) {
	group := app.Fiber.Group("/style-playground")

	group.Get("/:page", func(c *fiber.Ctx) error {

		ConditionalPage := map[string]templ.Component{
			"/":         view.ContentA(),
			"content-a": view.ContentA(),
			"content-b": view.ContentB(),
		}

		sidebarContent := map[string]interface{}{
			"Content A": "/content-a",
			"Content B": "/content-b",
		}

		renderer := internalApp.NewViewRenderer(c, view.Playground).ConditionalPage("/style-playground", ConditionalPage)
		return renderer.SetLocals(map[string]interface{}{
			"sidebar_content": sidebarContent,
		}).Render()
	})
}
