package handler

import (
	"github.com/Michael-Sjogren/gotempl/internal/views/pages"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct{}

func Render(c *fiber.Ctx, comp templ.Component) error {
	c.Response().Header.Add("Content-Type", "text/html")
	return comp.Render(c.Context(), c.Response().BodyWriter())
}

func (h *HomeHandler) HandlerHomePageView(c *fiber.Ctx) error {
	return Render(c, pages.HomePage())
}
