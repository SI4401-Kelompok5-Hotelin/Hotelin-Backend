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

	var hotel models.Hotel

	err := database.DB.Where("id = ?", req.HotelID).Find(&hotel).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to create booking",
		})
	}
	
	var room models.Room

	err = database.DB.Where("id = ?", req.RoomID).Find(&room).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to create booking",
		})
	}

	var additionalPrice float64

	if req.Covid == true {
		additionalPrice = 31000
	} else if req.Covid == false {
		additionalPrice = 0
	} else {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	var totalPrice float64

	totalPrice += (room.Price * float64(req.Duration)) + additionalPrice

	booking := models.Booking{
		UserID:			user.UserID,
		HotelID:		req.HotelID,
		HotelName:	hotel.Name,
		RoomID:			req.RoomID,
		RoomName:		room.Name,
		Covid:			req.Covid,
		TotalPrice:	totalPrice,
		CheckIn:		req.CheckIn,
		CheckOut:		req.CheckOut,
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
		HotelID:		hotel.ID,
		HotelName:	hotel.Name,
		RoomName:		room.Name,
		Duration:		booking.Duration,
		TotalPrice:	booking.TotalPrice,
		CheckIn:		booking.CheckIn,
		CheckOut:		booking.CheckOut,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Booking created",
		"data":    bookingResponse,
	})
}

func DeleteBooking(c *fiber.Ctx) error {
	id := c.Params("id")

	var booking models.Booking

	err := database.DB.First(&booking, id).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	err = database.DB.Delete(&booking).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Booking deleted",
	})
}

func ShowAllBooking(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserToken)

	booking := []models.Booking{}

	err := database.DB.Where("user_id = ?", user.UserID).Find(&booking).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get cart list",
		})
	}

	bookingResponse := []models.BookingResponse{}

	for _, booking := range booking {
		bookingResponse = append(bookingResponse, models.BookingResponse{
		ID:					booking.ID,
		HotelID:		booking.HotelID,
		HotelName:	booking.HotelName,
		RoomName:		booking.RoomName,
		Duration:		booking.Duration,
		TotalPrice:	booking.TotalPrice,
		CheckIn:		booking.CheckIn,
		CheckOut:		booking.CheckOut,
		})
	}

	if len(bookingResponse) == 0 {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Booking is empty",
			"status":  "success",
			"data":    []string{},
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Booking found",
		"data":    bookingResponse,
	})

}