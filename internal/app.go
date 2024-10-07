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
	AppInstance := &app.AppInstance{Fiber: fiberApp, Service: service}
	module.ModuleLoader(AppInstance)

	fiberApp.Static("/assets", "./assets")

	log.Fatal(fiberApp.Listen(":3000"))
}
