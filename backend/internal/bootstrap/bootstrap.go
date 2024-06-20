package bootstrap

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/run-bigpig/skplayer/internal/config"
	"github.com/run-bigpig/skplayer/internal/route"
	"io/fs"
	"log"
	"net/http"
)

func Boot(web embed.FS) {
	config.Set()
	app := fiber.New(fiber.Config{DisableStartupMessage: false})
	subFS, err := fs.Sub(web, "web")
	if err != nil {
		log.Fatalln(err)
	}
	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(subFS),
	}))
	app.Use(func(c *fiber.Ctx) error {
		// 如果路由以 /api 开头，则继续处理请求
		if c.Path() == "/api" || len(c.Path()) > 4 && c.Path()[:4] == "/api" {
			return c.Next()
		}

		// 否则返回 index.html
		return c.SendFile("web/index.html")
	})
	route.NewWeb(app)
	app.Listen(":8080")
}
