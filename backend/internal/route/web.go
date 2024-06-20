package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/skplayer/internal/controller"
)

func NewWeb(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/class", controller.Class)
	api.Post("/detail", controller.Detail)
	api.Post("/list", controller.List)
	api.Post("/search", controller.Search)
	api.Get("/getsource", controller.GetSourceList)
	api.Post("/setting", controller.Setting)
}
