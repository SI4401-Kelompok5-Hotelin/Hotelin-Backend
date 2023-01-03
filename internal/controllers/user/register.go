package user

import (
	"net/http"
	"Hotelin-BE/internal/database"
	"Hotelin-BE/internal/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	req := models.UserRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	var userType models.Type

	err = database.DB.Where("name = ?", "user").First(&userType).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	newUser := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Address:  req.Address,
		Password: string(hashedPassword),
		Type:     userType.Name,
	}

	// check if user already exists
	var existingUser models.User
	err = database.DB.Where("email = ?", newUser.Email).First(&existingUser).Error
	if err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
		})
	}


	err = database.DB.Create(&newUser).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	newUserType := models.UserType{
		UserID: newUser.ID,
		TypeID: userType.ID,
	}

	err = database.DB.Create(&newUserType).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
	})



}
