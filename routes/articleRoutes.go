package routes

import (
	"github.com/Micxxo/artikel-vision-be/controllers/articleController"
	"github.com/gofiber/fiber/v2"
)

func ArticleRoutes(api fiber.Router) {
	article := api.Group("/article")

	article.Get("/:limit/:page", articleController.Index)
	article.Get("/:id", articleController.Show)
	article.Post("/", articleController.Create)
	article.Patch("/:id", articleController.Update)
	article.Delete("/:id", articleController.Delete)
}
