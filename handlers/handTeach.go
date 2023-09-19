package handlers

import (
	"fmt"
	"school/database"
	"school/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Add new teacher
func TeachAdd(c *fiber.Ctx) error {
	query := models.Teacher{}
	if err := c.BodyParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	queryCheck := new(models.Teacher)
	//check ID if exists
	if gormdb := database.DB.Db.Where("id = ?", query.ID).First(&queryCheck); gormdb.RowsAffected != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "teacher with this ID already exists",
			"teacher": queryCheck,
		})
	}
	//check subject if exists
	checkSample := new(models.Subject)
	if gormdb := database.DB.Db.Where("name = ?", query.Subject).First(&checkSample); gormdb.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("subject %s does not exists in our school", query.Subject),
		})
	}
	database.DB.Db.Create(&query)
	return c.Status(fiber.StatusOK).JSON(query)
}

// Show All
func TeachShow(c *fiber.Ctx) error {
	query := []models.Teacher{}
	if gormdb := database.DB.Db.Order("id asc").Find(&query); gormdb.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "there are no teachers in our school",
		})
	}
	return c.Status(fiber.StatusOK).JSON(query)
}

// Show by ID
func TeachShowID(c *fiber.Ctx) error {
	needID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad params request",
		})
	}
	query := new(models.Teacher)
	if gormdb := database.DB.Db.Where("id = ?", needID).First(&query); gormdb.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("teacher with id=%d does not exist", needID),
		})
	}
	return c.Status(fiber.StatusOK).JSON(query)
}

// Show by subject
func TeachShowSub(c *fiber.Ctx) error {
	query := new(models.Subject)
	if err := c.BodyParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	result := []models.Teacher{}
	if gormdb := database.DB.Db.Where("subject = ?", query.Name).Order("id asc").Find(&result); gormdb.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("there is noone studying %s", query.Name),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

// Delete Teacher
func TeachDelete(c *fiber.Ctx) error {
	needID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad params request",
		})
	}
	teachSample := models.Teacher{}
	if gormdb := database.DB.Db.Where("id = ?", needID).Delete(&teachSample); gormdb.RowsAffected == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": fmt.Sprintf("there is noone with id=%d", needID),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "teacher deleted successfully",
	})
}

// Update Teacher by ID
func TeachUpdate(c *fiber.Ctx) error {
	needID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad params request",
		})
	}
	OldTeach := new(models.Teacher)
	if dbgorm := database.DB.Db.Where("id = ?", needID).First(&OldTeach); dbgorm.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("there is noone with id=%d", needID),
		})
	}
	teacher := new(models.Teacher)
	if err := c.BodyParser(&teacher); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	if teacher.Name != "" {
		database.DB.Db.Model(&OldTeach).Where("id = ?", needID).Update("name", teacher.Name)
	}
	if teacher.Surname != "" {
		database.DB.Db.Model(&OldTeach).Where("id = ?", needID).Update("surname", teacher.Surname)
	}
	if teacher.Grade != 0 {
		database.DB.Db.Model(&OldTeach).Where("id = ?", needID).Update("grade", teacher.Grade)
	}
	if teacher.Subject != "" {
		checkSample := new(models.Subject)
		if gormdb := database.DB.Db.Where("name = ?", teacher.Subject).First(&checkSample); gormdb.RowsAffected == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message":     fmt.Sprintf("subject %s does not exists in our school", teacher.Subject),
				"new teacher": OldTeach,
			})
		}
		database.DB.Db.Model(&OldTeach).Where("id = ?", needID).Update("subject", teacher.Subject)
	}
	return c.Status(fiber.StatusOK).JSON(OldTeach)
}
