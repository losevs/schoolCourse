package handlers

import (
	"school/database"
	"school/models"

	"github.com/gofiber/fiber/v2"
)

func DirAdd(c *fiber.Ctx) error {
	sampCheck := models.Director{}
	if gormdb := database.DB.Db.First(&sampCheck); gormdb.RowsAffected != 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":  "director already exists",
			"director": sampCheck,
		})
	}
	query := models.Director{}
	if err := c.BodyParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	database.DB.Db.Create(&query)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "new director updated!",
	})
}

func ShowDir(c *fiber.Ctx) error {
	DirSamp := models.Director{}
	if gormdb := database.DB.Db.First(&DirSamp); gormdb.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "there is no director",
		})
	}
	return c.Status(fiber.StatusOK).JSON(DirSamp)
}

func DeleteDir(c *fiber.Ctx) error {
	DirSamp := models.Director{}
	if gormdb := database.DB.Db.First(&DirSamp); gormdb.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "there is no director",
		})
	}
	database.DB.Db.Where("name = ?", DirSamp.Name).Delete(&DirSamp)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "director deleted successfully",
	})
}
