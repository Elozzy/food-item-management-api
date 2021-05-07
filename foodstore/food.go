package foodstore

import (
	"go_restapi_fiber/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name     string `json: "name"`
	Price    int    `json: "price"`
	Quantity int    `json: "quantity"`
}

func GetFoodItems(c *fiber.Ctx) error {
	db := database.DBConn
	var foodItems []Food
	db.Find(&foodItems)
	return c.JSON(foodItems)
}

func GetFoodItem(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var foodItem Food
	db.Find(&foodItem, id)
	return c.JSON(foodItem)
}

func NewFoodItem(c *fiber.Ctx) error {
	db := database.DBConn

	foodItem := new(Food)
	if err := c.BodyParser(foodItem); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&foodItem)
	return c.JSON(foodItem)
}

func DeleteFoodItem(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var foodItem Food
	db.First(&foodItem, id)
	if foodItem.Name == " " {
		return c.Status(500).SendString("No food item found with the given ID")
	}
	db.Delete(&foodItem)
	return c.SendString("Food item successfully deleted")
}


