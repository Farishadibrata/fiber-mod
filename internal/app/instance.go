package app

import (
	"context"

	"farishadibrata.com/fiber-modular/internal/services"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

type AppInstance struct {
	Fiber   *fiber.App
	Service *services.Service
}

type Renderer struct {
	Template            func(contents templ.Component) templ.Component
	Page                templ.Component
	Fiber               *fiber.Ctx
	ConditionalPagePath string
}

func NewViewRenderer(c *fiber.Ctx, template func(contents templ.Component) templ.Component) *Renderer {
	renderer := &Renderer{Template: template, Fiber: c}
	return renderer
}

// Use this for single page / non tabs.
func (r *Renderer) SetPage(page templ.Component) *Renderer {
	r.Page = page
	return r
}

// Use this for multiple rendering with /* path
func (r *Renderer) ConditionalPage(path string, pages map[string]templ.Component, notFoundHandler templ.Component) *Renderer {
	r.ConditionalPagePath = path
	r.Fiber.Locals("CURRENT_PATH", path)
	currentPath := r.Fiber.Params("page")
	if pages[currentPath] == nil {
		r.Page = notFoundHandler
		return r
	}
	r.Page = pages[currentPath]

	return r
}

// And use this to render another link that is not the current page.
func URLDynamic(ctx context.Context, link string) templ.SafeURL {
	return templ.URL(ctx.Value("CURRENT_PATH").(string) + "/" + link)
}

func (r *Renderer) SetLocals(vars map[string]interface{}) *Renderer {
	if r.ConditionalPagePath != "" {
		r.Fiber.Locals("CURRENT_PATH", r.ConditionalPagePath)
	}
	for key, value := range vars {

		r.Fiber.Locals(key, value)
	}
	return r
}

func (r *Renderer) Render() error {
	r.Fiber.Set("Content-Type", "text/html")

	header := r.Fiber.GetReqHeaders()
	partialRender := false
	for key, value := range header {
		if key == "Hx-Request" {
			if value[0] == "true" {
				partialRender = true
			}
		}
	}

	if partialRender {
		r.Fiber.Set("HX-Trigger-After-Swap", "reInitialize")
		return r.Page.Render(r.Fiber.Context(), r.Fiber.Response().BodyWriter())
	}

	return r.Template(r.Page).Render(r.Fiber.Context(), r.Fiber.Response().BodyWriter())
}
