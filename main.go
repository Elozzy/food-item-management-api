package main

import (
	"fmt"
	"go_restapi_fiber/database"
	"go_restapi_fiber/foodstore"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	_"github.com/mattn/go-sqlite3"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World !")
}

//Routes
func setupRoutes(app *fiber.App) error {
	app.Get("/api/v1/food", foodstore.GetFoodItems)
	app.Get("/api/v1/food/:id", foodstore.GetFoodItem)
	app.Post("/api/v1/food", foodstore.NewFoodItem)
	app.Delete("/api/v1/food/:id", foodstore.DeleteFoodItem)
	return nil
}

//initialize database
func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open ("fooditems.db"))
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")

	database.DBConn.AutoMigrate(&foodstore.Food{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)

	app.Listen(":3000")
}
