package usercontroller

import (
	"belajar-fiber-gorm/database"
	"belajar-fiber-gorm/models"
	"belajar-fiber-gorm/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": users,
	})
}

func Show(c *fiber.Ctx) error {

	id := c.Params("id")

	var user models.User
	if database.DB.First(&user, id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": user,
	})
}

func Create(c *fiber.Ctx) error {

	user := new(models.UserCreate)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errValidate.Error(),
		})
	}

	newUser := models.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	newUser.Password = hashedPassword

	errCreate := database.DB.Create(&newUser).Error
	if errCreate != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errCreate.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": newUser,
	})
}

func Update(c *fiber.Ctx) error {
	// check body
	userReq := new(models.UserCreate)

	if err := c.BodyParser(userReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// check data
	id := c.Params("id")

	var user models.User
	if database.DB.First(&user, id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	// update data
	if userReq.Name != "" {
		user.Name = userReq.Name
	}
	if userReq.Address != "" {
		user.Address = userReq.Address
	}
	if userReq.Email != "" {
		user.Email = userReq.Email
	}
	if userReq.Phone != "" {
		user.Phone = userReq.Phone
	}

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errUpdate.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": user,
	})
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	err := database.DB.Debug().First(&user, id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errDelete := database.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errDelete.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "user deleted",
	})
}
