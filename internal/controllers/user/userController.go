package user

import (
	"net/http"
	"Hotelin-BE/internal/database"
	"Hotelin-BE/internal/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func UserDetail(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserToken)

	var userDetail models.User

	err := database.DB.Where("id = ?", user.UserID).First(&userDetail).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	userResponse := models.UserDetail{
		ID:    userDetail.ID,
		Name:  userDetail.Name,
		Email: userDetail.Email,
		Phone: userDetail.Phone,
		Address: userDetail.Address,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": userResponse,
	})
}

func ChangeProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserToken)

	req := models.User{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	
	var userDetail models.User
	
	err := database.DB.Where("id = ?", user.UserID).First(&userDetail).Error
	
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	
	userDetail.Name = req.Name
	userDetail.Phone = req.Phone
	userDetail.Email = req.Email
	userDetail.Address = req.Address
	userDetail.Password = string(hashedPassword)

	err = database.DB.Save(&userDetail).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success change profile",
	})

}