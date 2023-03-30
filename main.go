package main

import (
	"fiber-practice/database"
	"fiber-practice/handlers"
	"fiber-practice/models"
	"github.com/gofiber/fiber/v2"
)

func dbInit() { //connect db & make tables
	database.Connect()
	db := database.Get()

	// 테이블 자동 생성
	db.Migrator().DropTable(&models.User{}, &models.CreditCard{})
	db.AutoMigrate(&models.User{}, &models.CreditCard{})
}

func setUpRoutes(app *fiber.App) { //set api url to handler
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/user/:id", handlers.UserGet)
	app.Delete("/user/:id", handlers.UserDelete)
	app.Put("/user/:id", handlers.UserUpdate)
	app.Post("/user", handlers.UserCreate)
}

func main() {
	dbInit()
	app := fiber.New()
	setUpRoutes(app)
	app.Listen(":3000")
}
