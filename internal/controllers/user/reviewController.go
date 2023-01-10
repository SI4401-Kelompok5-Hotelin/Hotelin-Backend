package user

import (
	"Hotelin-BE/internal/database"
	"Hotelin-BE/internal/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateReview(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserToken)

	req := models.ReviewRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	review := models.Review{}

	err := database.DB.Where("id = ?", user.UserID).First(&review).Error

	if err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Review sudah ada",
		})
	}

	username := models.User{}

	err = database.DB.Where("id = ?", user.UserID).Find(&username).Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	newReview := models.Review{
		UserID:			user.UserID,
		UserName:		username.Name,
		HotelID:		req.HotelID,
		Review:			req.Review,
	}

	err = database.DB.Create(&newReview).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	response := models.ReviewResponse{
		ID:				newReview.ID,
		UserName:	username.Name,
		Review:		newReview.Review,
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Review added",
		"data":			response,
	})
}

func GetReview(c *fiber.Ctx) error {

	review := []models.Review{}

	id := c.Params("id")


	err := database.DB.Where("hotel_id = ?", id).Find(&review).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	reviewResponse := []models.ReviewResponse{}

	for _, review := range review {
		reviewResponse = append(reviewResponse, models.ReviewResponse{
			ID:      				review.ID,
			UserName:    		review.UserName,
			Review:   			review.Review,
		})
	}

	

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Review found",
		"data":    reviewResponse,
	})
}