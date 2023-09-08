package main

import (
	"fmt"
	"school/database"
	"school/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("err: ", err)
	}
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)
	app.Listen(":80") //localhost:80
}

func setupRoutes(app *fiber.App) {
	app.Get("/show", handlers.Show)
	app.Get("/show/class/:grade", handlers.ShowGrade)
	app.Get("/show/:id", handlers.ShowExact)

	app.Post("/new", handlers.Create)

	app.Delete("/delete/:id", handlers.Delete)

	app.Patch("/update/:id", handlers.Update)
}