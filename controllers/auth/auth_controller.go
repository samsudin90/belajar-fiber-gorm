package auth

import (
	"belajar-fiber-gorm/database"
	"belajar-fiber-gorm/models"
	"belajar-fiber-gorm/utils"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(models.LoginRequest)

	// check body json
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// validate
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errValidate.Error(),
		})
	}

	// check user
	var user models.User
	err := database.DB.First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// check password
	isValid := utils.CheckPassword(loginRequest.Password, user.Password)
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "email or password wrong",
		})
	}

	// generate jwt
	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	// untuk expire token
	// claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token, errToken := utils.GenerateToken(&claims)
	if errToken != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errToken.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("foto")
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	filename := file.Filename

	errSave := c.SaveFile(file, fmt.Sprintf("./public/assets/%s", filename))

	if errSave != nil {
		return c.JSON(fiber.Map{
			"message": errSave.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "oke",
	})
}
