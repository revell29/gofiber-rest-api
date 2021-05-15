package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	itemhandler "github.com/revell29/gofiber-rest-api/modules"
)

func InitRoute(app *fiber.App) {
	api := app.Group("/api", logger.New())  
	
	//routes
	api.Get("/items", itemhandler.GetAllItems)
	api.Get("/items/:id", itemhandler.DetailItems)
}