package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"time"
)

type UserBooking struct {
		ID				string	`json:"id" gorm:"primary_key, type:uid, default:uuid_generate_v4()"`
		UserID		string	`json:"user_id"`
		User			User		`json:"user" gorm:"foreignKey:UserID"`
		HotelID		string	`json:"hotel_id"`
		Hotel			Hotel		`json:"hotel" gorm:"foreignKey:HotelID"`
		RoomID		string	`json:"room_id"`
		Room			Room		`json:"room" gorm:"foreignKey:RoomID"`
		BookingID	string	`json:"booking_id"`
		Booking		Booking	`json:"booking" gorm:"foreignKey:BookingID"`
		Duration	string	`json:"duration"`
		CheckIn		time.Time	`json:"check_in"`
		CheckOut	time.Time	`json:"check_out"`
}

func (u *UserBooking) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}

type UserBookingRequest struct {
		HotelID			string	`json:"hotel_id"`
		RoomID			string	`json:"room_id"`
		Duration		string	`json:"duration"`
}