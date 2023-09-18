package handlers

import (
	"fmt"
	"school/database"
	"school/models"

	"github.com/gofiber/fiber/v2"
)

func AddSub(c *fiber.Ctx) error {
	query := models.Subject{}
	if err := c.BodyParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	sampleCheck := models.Subject{}
	if dbgorm := database.DB.Db.Where("name = ?", query.Name).First(&sampleCheck); dbgorm.RowsAffected != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "this subject already exists",
		})
	}
	database.DB.Db.Create(&query)
	return c.Status(fiber.StatusOK).JSON(query)
}

func ShowSubs(c *fiber.Ctx) error {
	sample := []models.Subject{}
	if dbgorm := database.DB.Db.Find(&sample); dbgorm.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "there are no subjects",
		})
	}
	return c.Status(fiber.StatusOK).JSON(sample)
}

func DelSub(c *fiber.Ctx) error {
	query := models.Subject{}
	if err := c.BodyParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	if gormdb := database.DB.Db.Where("name = ?", query.Name).Delete(&query); gormdb.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("subject %s does not exists", query.Name),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "subject deleted successfully",
	})
}
