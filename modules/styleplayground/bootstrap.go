package styleplayground

import (
	internalApp "farishadibrata.com/fiber-modular/internal/app"
	"farishadibrata.com/fiber-modular/modules/styleplayground/models"
	"farishadibrata.com/fiber-modular/modules/styleplayground/utils"
	view "farishadibrata.com/fiber-modular/modules/styleplayground/view"
	"github.com/gofiber/fiber/v2"
)

func Bootstrap(app *internalApp.AppInstance) {
	group := app.Fiber.Group("/style-playground")

	playgroundHandler := func(c *fiber.Ctx) error {

		sidebarContent := []models.Sidebar{
			{
				Title:   "Home",
				Path:    "",
				Content: view.Home(),
			},
			{
				Title:   "Button",
				Path:    "button",
				Content: view.Button(),
			},
			{
				Title:   "Alert",
				Path:    "alert",
				Content: view.Alert(),
			},
		}

		pages := utils.ToMapStringInterface(sidebarContent)
		renderer := internalApp.NewViewRenderer(c, view.Playground).ConditionalPage("/style-playground", pages, view.Home())
		return renderer.SetLocals(map[string]interface{}{
			"sidebar_content": sidebarContent,
		}).Render()
	}
	group.Get("/:page", playgroundHandler)
	group.Get("/", playgroundHandler)
}
