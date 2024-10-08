package internal

import (
	"log"

	"farishadibrata.com/fiber-modular/internal/app"
	"farishadibrata.com/fiber-modular/internal/module"
	"farishadibrata.com/fiber-modular/internal/services"
	"github.com/gofiber/fiber/v2"
)

func NewApp() {
	fiberApp := fiber.New()
	service := services.NewService()
	fiberApp.Static("/assets", "./assets")
	fiberApp.Static("/x", ".")

	AppInstance := &app.AppInstance{Fiber: fiberApp, Service: service}
	module.ModuleLoader(AppInstance)

	log.Fatal(fiberApp.Listen(":3000"))
}
