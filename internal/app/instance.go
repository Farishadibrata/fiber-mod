package app

import (
	"farishadibrata.com/fiber-modular/internal/services"
	"github.com/Oudwins/tailwind-merge-go/pkg/twmerge"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

type IRedirect interface {
	Redirect(renderer *Renderer) error
}

type RedirectTo struct {
	To string
}

func (r *RedirectTo) Redirect(renderer *Renderer) error {
	return renderer.Fiber.Redirect(r.To)
}

type RedirectPage struct {
	Page templ.Component
}

func (r *RedirectPage) Redirect(renderer *Renderer) error {
	renderer.Page = r.Page
	return nil
}

type AppInstance struct {
	Fiber   *fiber.App
	Service *services.Service
	TW      twmerge.TwMergeFn
}

type RedirectType string

const (
	REDIRECT_PAGE RedirectType = "REDIRECT_PAGE"
	REDIRECT_TO   RedirectType = "REDIRECT_TO"
)

type Renderer struct {
	Template            func(contents templ.Component) templ.Component
	NotFoundHandler     IRedirect
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
func (r *Renderer) PathPage(path string, pages map[string]templ.Component, NotFound IRedirect) *Renderer {
	r.ConditionalPagePath = path
	r.Fiber.Locals("CURRENT_PATH", path)
	currentPath := r.Fiber.Params("page")
	r.NotFoundHandler = NotFound
	r.Page = pages[currentPath]
	return r
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

	if r.Page == nil {
		return r.NotFoundHandler.Redirect(r)
	}
	if partialRender {
		return r.Page.Render(r.Fiber.Context(), r.Fiber.Response().BodyWriter())
	}

	return r.Template(r.Page).Render(r.Fiber.Context(), r.Fiber.Response().BodyWriter())
}
