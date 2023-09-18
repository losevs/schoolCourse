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
	//Subjects
	subs := app.Group("/subs")
	subs.Get("/show", handlers.ShowSubs)
	subs.Post("/add", handlers.AddSub)
	subs.Delete("/del", handlers.DelSub)

	//Director
	dir := app.Group("/dir")
	dir.Get("/show", handlers.ShowDir)
	dir.Post("/add", handlers.DirAdd)
	dir.Delete("/del", handlers.DeleteDir)

	//Teachers
	teach := app.Group("/teach")
	teachShow := teach.Group("/show")
	teachShow.Get("/", handlers.TeachShow)
	// teachShow.Get("/:id") // -------------------------
	// teachShow.Get("/sub/:subj") // ---------------------

	teach.Post("/new", handlers.TeachAdd)

	// teach.Delete("/delete/:id") // --------------------------

	// teach.Patch("/update/:id") // --------------------------

	//Kids
	show := app.Group("/show")
	show.Get("/", handlers.Show)
	show.Get("/class/:grade", handlers.ShowGrade)
	show.Get("/:id", handlers.ShowExact)

	app.Post("/new", handlers.Create)

	app.Delete("/delete/:id", handlers.Delete)

	app.Patch("/update/:id", handlers.Update)
}
