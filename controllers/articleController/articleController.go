package articleController

import (
	"github.com/Micxxo/artikel-vision-be/databases"
	"github.com/Micxxo/artikel-vision-be/helpers"
	"github.com/Micxxo/artikel-vision-be/models"
	"github.com/Micxxo/artikel-vision-be/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var validate = validator.New()

func Index(c *fiber.Ctx) error {
	var articles []models.Post

	pagination, err := utils.Paginate(c, databases.DB, &articles)
	if err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(pagination)
}

func Show(c *fiber.Ctx) error {
	id := c.Params("id")
	var articles models.Post

	if result := databases.DB.First(&articles, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return helpers.SendErrorResponse(c, fiber.StatusNotFound, result.Error.Error())
		}
		return helpers.SendErrorResponse(c, fiber.StatusInternalServerError, result.Error.Error())
	}

	return helpers.SendSuccessResponse(c, fiber.StatusOK, articles)
}

func Create(c *fiber.Ctx) error {
	var articles models.Post
	if err := c.BodyParser(&articles); err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	// Input validation
	if err := validate.Struct(articles); err != nil {
		errorMessages := helpers.MapValidationErrors(err)
		return helpers.SendErrorResponse(c, fiber.StatusBadRequest, errorMessages)
	}

	if err := databases.DB.Create(&articles).Error; err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SendSuccessResponse(c, fiber.StatusOK, articles)
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var articles models.Post

	if err := c.BodyParser(&articles); err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	// Input validation
	if err := validate.Struct(articles); err != nil {
		errorMessages := helpers.MapValidationErrors(err)
		return helpers.SendErrorResponse(c, fiber.StatusBadRequest, errorMessages)
	}

	if databases.DB.Where("id = ?", id).Updates(&articles).RowsAffected == 0 {
		return helpers.SendErrorResponse(c, fiber.StatusBadRequest, "Update article failed")
	}

	// Reload to get newest data
	databases.DB.First(&articles, id)

	return helpers.SendSuccessResponse(c, fiber.StatusOK, articles)
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	var articles models.Post
	if databases.DB.Delete(&articles, id).RowsAffected == 0 {
		return helpers.SendErrorResponse(c, fiber.StatusBadRequest, "Delete article failed")
	}

	return helpers.SendSuccessResponse(c, fiber.StatusOK, "Delete article success")
}
