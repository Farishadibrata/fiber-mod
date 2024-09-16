package internal

import (
	"log"

	"farishadibrata.com/fiber-modular/internal/app"
	"farishadibrata.com/fiber-modular/internal/module"
	"farishadibrata.com/fiber-modular/internal/services"
	"github.com/gofiber/fiber/v2"
)

func NewApp() {
	fiber := fiber.New()

	service := services.NewService()
	AppInstance := &app.AppInstance{Fiber: fiber, Service: service}
	module.ModuleLoader(AppInstance)

	fiber.Static("/assets", "./assets")
	log.Fatal(fiber.Listen(":3000"))
}
