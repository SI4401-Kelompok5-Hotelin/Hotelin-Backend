package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"time"
)

type Booking struct {
		ID					string		`json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
		UserID 			string		`json:"user_id"`
		User				User			`json:"user" gorm:"foreignKey:UserID"`
		HotelID			string		`json:"hotel_id"`
		Hotel				Hotel			`json:"hotel" gorm:"foreignKey:HotelID"`
		RoomID			string		`json:"room_id"`
		Room				Room			`json:"room" gorm:"foreignKey:RoomID"`
		Covid				string		`json:"covid"`
		TotalPrice	float64		`json:"total_price"`
		CheckIn			time.Time	`json:"check_in"`
		CheckOut		time.Time	`json:"check_out"`
		Duration		string		`json:"duration"`
		CreatedAt		time.Time	`json:"created_at"`
}

func (u *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}

type BookingRequest struct {
		HotelID		string		`json:"hotel_id"`
		RoomID		string		`json:"room_id"`
		Covid			string		`json:"covid"`
		Duration	string		`json:"duration"`
		CheckIn		time.Time	`json:"check_in"`
		CheckOut	time.Time	`json:"check_out"`
}

type BookingResponse struct {
		ID					string		`json:"id"`
		HotelName		string		`json:"hotel_name"`
		RoomName		string		`json:"room_name"`
		Duration		string		`json:"duration"`
		TotalPrice	float64		`json:"total_price"`
		CheckIn			time.Time	`json:"check_in"`
		CheckOut		time.Time	`json:"check_out"`
}