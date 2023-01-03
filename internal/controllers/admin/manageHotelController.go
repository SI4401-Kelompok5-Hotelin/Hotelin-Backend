package admin

import (
	"net/http"

	"Hotelin-BE/internal/models"
	"Hotelin-BE/internal/database"
	"github.com/gofiber/fiber/v2"
)

func ShowAllHotel(c *fiber.Ctx) error {
	hotel := []models.Hotel{}

	// show all hotel in database with owner and don't show owner password
	err := database.DB.Find(&hotel).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get hotel",
		})
	}

	hotelResponse := []models.HotelResponse{}

	for _, hotel := range hotel {
		hotelResponse = append(hotelResponse, models.HotelResponse{
			ID:      hotel.ID,
			Name:    hotel.Name,
			Email:   hotel.Email,
			Phone:   hotel.Phone,
			Address: hotel.Address,
			Rating:	 hotel.Rating,
		})
	}

	

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "hotel found",
		"data":    hotelResponse,
	})

}