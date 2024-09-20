package internal

import (
	"log"

	"farishadibrata.com/fiber-modular/internal/app"
	"farishadibrata.com/fiber-modular/internal/module"
	"farishadibrata.com/fiber-modular/internal/services"
	"github.com/gofiber/fiber/v2"

	twmerge "github.com/Oudwins/tailwind-merge-go/pkg/twmerge"
)

var TwMerger twmerge.TwMergeFn

func NewApp() {
	fiber := fiber.New()

	service := services.NewService()
	AppInstance := &app.AppInstance{Fiber: fiber, Service: service, TW: TwMerger}
	module.ModuleLoader(AppInstance)

	config := twmerge.MakeDefaultConfig()

	// do your modifications here

	// create the merger
	TwMerger = twmerge.CreateTwMerge(config, nil)

	fiber.Static("/assets", "./assets")
	log.Fatal(fiber.Listen(":3000"))
}
