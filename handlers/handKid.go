package handlers

import (
	"fmt"
	"school/database"
	"school/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Add new kid to a database
func Create(c *fiber.Ctx) error {
	query := new(models.Kid)
	if err := c.BodyParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad request",
		})
	}
	queryCheckID := new(models.Kid)
	if gormdb := database.DB.Db.Where("id = ?", query.ID).First(queryCheckID); gormdb.RowsAffected != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "this ID already exists",
			"kid":     queryCheckID,
		})
	}
	database.DB.Db.Create(&query)
	return c.Status(fiber.StatusOK).JSON(query)
}

// Show every kid
func Show(c *fiber.Ctx) error {
	sample := []models.Kid{}
	if dbGorm := database.DB.Db.Order("id asc").Find(&sample); dbGorm.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "empty",
		})
	}
	return c.Status(fiber.StatusOK).JSON(sample)
}

// Show kids from given grade
func ShowGrade(c *fiber.Ctx) error {
	gradeUint64, err := strconv.ParseUint(c.Params("grade"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	sample := []models.Kid{}
	if dbGorm := database.DB.Db.Where("grade = ?", gradeUint64).Order("id asc").Find(&sample); dbGorm.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Class %d is empty", gradeUint64),
		})
	}
	return c.Status(fiber.StatusOK).JSON(sample)
}

// Show exact kid by ID
func ShowExact(c *fiber.Ctx) error {
	needID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad params request",
		})
	}
	sample := new(models.Kid)
	if dbgorm := database.DB.Db.Where("id = ?", needID).First(&sample); dbgorm.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "there is no kid with that id",
		})
	}
	return c.Status(fiber.StatusOK).JSON(sample)
}

// Delete kiddo
func Delete(c *fiber.Ctx) error {
	needID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	if gormdb := database.DB.Db.Where("id = ?", needID).Delete(&models.Kid{}); gormdb.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "there is no kids with this id",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "kid deleted succsessfully!",
	})
}

// Update kid w/out id
func Update(c *fiber.Ctx) error {
	needID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	Kiddo := new(models.Kid)
	if gormdb := database.DB.Db.Where("id = ?", needID).First(&Kiddo); gormdb.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "there is no kid with this ID",
		})
	}
	query := new(models.Kid)
	if err := c.BodyParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	if query.Name != "" {
		database.DB.Db.Model(&Kiddo).Where("id = ?", needID).Update("name", query.Name)
	}
	if query.Surname != "" {
		database.DB.Db.Model(&Kiddo).Where("id = ?", needID).Update("surname", query.Surname)
	}
	if query.Age != 0 {
		database.DB.Db.Model(&Kiddo).Where("id = ?", needID).Update("age", query.Age)
	}
	if query.Grade != 0 {
		database.DB.Db.Model(&Kiddo).Where("id = ?", needID).Update("grade", query.Grade)
	}
	return c.Status(fiber.StatusOK).JSON(Kiddo)
}
