package user

import (
	"Hotelin-BE/internal/database"
	"Hotelin-BE/internal/models"
	"net/http"
	"github.com/gofiber/fiber/v2"
)

func TopUpBalance(c *fiber.Ctx) error {
	req := models.UserBalanceRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	user := c.Locals("user").(models.UserToken)

	balance := models.UserBalance{}

	err := database.DB.Where("user_id = ?", user.UserID).First(&balance).Error
	if err != nil {
		// if user id is not found, create new user balance
		balance = models.UserBalance{
			UserID:  user.UserID,
			Balance: req.Balance,
		}

		err = database.DB.Create(&balance).Error

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Failed to create balance",
			})
		}
	} else {
		// if user id is found, update balance
		balance.Balance += req.Balance

		err = database.DB.Save(&balance).Error

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Failed to update balance",
			})
		}
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success topup balance",
	})

}

func GetBalance(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserToken)

	var userBalance models.UserBalance

	err := database.DB.Where("user_id = ?", user.UserID).First(&userBalance).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"balance": userBalance.Balance,
	})
}