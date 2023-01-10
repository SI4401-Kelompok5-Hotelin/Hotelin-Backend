package hotel

import (
	"Hotelin-BE/internal/database"
	"Hotelin-BE/internal/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoom(c *fiber.Ctx) error {
	// Get userID from JWT token
	user := c.Locals("user").(models.UserToken)

	req := models.RoomRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	rooms := models.Room{}

	err := database.DB.Where("id = ?", user.UserID).First(&rooms).Error

	if err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Room already registered",
		})
	}

	newRoom := models.Room{
		Name:    			 req.Name,
		Price:   			 req.Price,
		Description:   req.Description,
		HotelID: 			 req.HotelID,
	}

	err = database.DB.Create(&newRoom).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	response := models.RoomResponse{
		ID:      			 newRoom.ID,
		Name:    			 newRoom.Name,
		Price:   			 newRoom.Price,
		Description:   newRoom.Description,
		HotelID: 			 newRoom.HotelID,
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Room registered",
		"data":			response,
	})

}

func UpdateRoom(c *fiber.Ctx) error {
	id := c.Params("id")

	var room models.Room

	err := database.DB.First(&room, id).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error id",
		})
	}

	req := models.RoomRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	err = database.DB.Model(&room).Updates(req).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Room updated",
	})
}

func DeleteRoom(c *fiber.Ctx) error {
	id := c.Params("id")

	var room models.Room

	err := database.DB.First(&room, id).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	err = database.DB.Delete(&room).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Room deleted",
	})
}

func GetRoomByHotelID(c *fiber.Ctx) error {
	room := []models.Room{}

	id := c.Query("id")

	err := database.DB.Where("hotel_id = ?", id).Find(&room).Error

	// show all room in database with owner and don't show owner password
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get room",
		})
	}

	roomResponse := []models.RoomResponse{}

	for _, room := range room {
		roomResponse = append(roomResponse, models.RoomResponse{
			ID:      				room.ID,
			Name:    				room.Name,
			Price:   				room.Price,
			Description:  	room.Description,
			HotelID: 				room.HotelID,
		})
	}

	

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "room found",
		"data":    roomResponse,
	})

}