package index

import (
	. "farishadibrata.com/fiber-modular/internal/app"
	"farishadibrata.com/fiber-modular/modules/styleplayground/models"
	. "farishadibrata.com/fiber-modular/modules/styleplayground/routes/alert/view"
	button "farishadibrata.com/fiber-modular/modules/styleplayground/routes/button/view"
	card "farishadibrata.com/fiber-modular/modules/styleplayground/routes/card/view"
	"farishadibrata.com/fiber-modular/modules/styleplayground/routes/index/view"
	input "farishadibrata.com/fiber-modular/modules/styleplayground/routes/input/view"
	tooltip "farishadibrata.com/fiber-modular/modules/styleplayground/routes/tooltip/view"
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
			Content: button.ButtonPage(),
		},
		{
			Title:   "Alert",
			Path:    "alert",
			Content: AlertPage(),
		},
		{
			Title:   "Card",
			Path:    "card",
			Content: card.CardPage(),
		},
		{
			Title:   "Input",
			Path:    "input",
			Content: input.InputPage(),
		},
		{
			Title:   "tooltip",
			Path:    "tooltip",
			Content: tooltip.TooltipPage(),
		},
	}

	pages := utils.ToMapStringInterface(sidebarContent)
	renderer := NewViewRenderer(c, view.Playground).ConditionalPage("/style-playground", pages, view.Home())

	return renderer.SetLocals(map[string]interface{}{
		"sidebar_content": sidebarContent,
	}).Render()
}
