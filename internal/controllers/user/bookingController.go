package user

import (
	"Hotelin-BE/internal/database"
	"Hotelin-BE/internal/models"
	"net/http"
	"github.com/gofiber/fiber/v2"
)

func CreateBooking(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserToken)

	req := models.BookingRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	var room models.Room
	var hotel models.Hotel

	var additionalPrice float64

	if req.Covid == "Yes" {
		additionalPrice = 31000
	} else if req.Covid == "No" {
		additionalPrice = 0
	} else {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	var totalPrice float64

	for _, roomPrice := range room {
		totalPrice += roomPrice.TotalPrice + additionalPrice
	}

	booking := models.Booking{
						UserID:			user.UserID,
						HotelID:		req.HotelID,
						RoomID:			req.RoomID,
						Covid:			req.Covid,
						TotalPrice:	totalPrice,
						CheckIn:		req.CheckIn,
						CheckOut		req.CheckOut,
						Duration:		req.Duration,
	}

	var userBalance models.UserBalance

	err = database.DB.Where("user_id = ?", user.UserID).Find(&userBalance).Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to create order Your balance is not enough",
		})
	}

	if userBalance.Balance < totalPrice {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to create order Your balance is not enough",
		})
	}

	// if sucess then create order
	err = database.DB.Create(&booking).Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to create order",
		})
	}

	userBalance.Balance -= totalPrice

	err = database.DB.Save(&userBalance).Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to create order",
		})
	}

	bookingResponse := models.BookingResponse{
		ID:					booking.ID,
		HotelName:	booking.
	}
}