package styleplayground

import (
	internalApp "farishadibrata.com/fiber-modular/internal/app"
	view "farishadibrata.com/fiber-modular/modules/styleplayground/view"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func Bootstrap(app *internalApp.AppInstance) {
	group := app.Fiber.Group("/style-playground")

	pageHandler := func(c *fiber.Ctx) error {

		ConditionalPage := map[string]templ.Component{
			"":          view.Index(),
			"content-a": view.ContentA(),
			"content-b": view.ContentB(),
		}

		sidebarContent := map[string]interface{}{
			"Content A": "/content-a",
			"Content B": "/content-b",
		}

		view := internalApp.NewViewRenderer(c, view.Playground)
		view.PathPage("/style-playground", ConditionalPage, &internalApp.RedirectTo{To: "content-a"})

		view.SetLocals(map[string]interface{}{
			"sidebar_content": sidebarContent,
		})

		return view.Render()
	}

	group.Get("/:page", pageHandler)
	group.Get("/", pageHandler)
}
