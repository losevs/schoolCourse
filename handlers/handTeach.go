package handlers

import (
	"fmt"
	"school/database"
	"school/models"

	"github.com/gofiber/fiber/v2"
)

func TeachAdd(c *fiber.Ctx) error {
	query := models.Teacher{}
	if err := c.BodyParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	queryCheck := new(models.Teacher)
	//check ID if exists
	if gormdb := database.DB.Db.Where("id = ?", query.ID).First(queryCheck); gormdb.RowsAffected != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "teacher with this ID already exists",
			"teacher": queryCheck,
		})
	}
	//check subject if exists
	checkSample := new(models.Subject)
	if gormdb := database.DB.Db.Where("name = ?", query.Subject).First(checkSample); gormdb.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("subject %s does not exists in our school", query.Subject),
		})
	}
	database.DB.Db.Create(&query)
	return c.Status(fiber.StatusOK).JSON(query)
}

func TeachShow(c *fiber.Ctx) error {
	query := []models.Teacher{}
	if gormdb := database.DB.Db.Find(&query); gormdb.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "there are no teachers in our school",
		})
	}
	return c.Status(fiber.StatusOK).JSON(query)
}
