package index

import (
	internalApp "farishadibrata.com/fiber-modular/internal/app"
	"farishadibrata.com/fiber-modular/modules/styleplayground/models"
	alertView "farishadibrata.com/fiber-modular/modules/styleplayground/routes/alert/view"
	"farishadibrata.com/fiber-modular/modules/styleplayground/routes/index/view"
	"farishadibrata.com/fiber-modular/modules/styleplayground/utils"
	"github.com/gofiber/fiber/v2"
)

var PlaygroundHandler = func(c *fiber.Ctx) error {

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
			Content: alertView.AlertPage(),
		},
	}

	pages := utils.ToMapStringInterface(sidebarContent)
	renderer := internalApp.NewViewRenderer(c, view.Playground).ConditionalPage("/style-playground", pages, view.Home())

	return renderer.SetLocals(map[string]interface{}{
		"sidebar_content": sidebarContent,
	}).Render()
}
