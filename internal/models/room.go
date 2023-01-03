package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Room struct {
	ID      		string      `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name    		string      `json:"name"`
	Price   		int64     	`json:"price"`
	Description string			`json:"description"`
	HotelID 		string      `json:"hotel_id"`
	Hotel   		Hotel 			`json:"hotel" gorm:"foreignKey:HotelID"`
}

func (r *Room) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.NewString()
	return
}

type RoomRequest struct {
	Name    		string      `json:"name"`
	Price   		int64     	`json:"price"`
	Description string			`json:"description"`
	HotelID 		string      `json:"hotel_id"`
}

type RoomResponse struct {
	ID					string			`json:"id"`
	Name    		string      `json:"name"`
	Price   		int64     	`json:"price"`
	Description string			`json:"description"`
	HotelID 		string      `json:"hotel_id"`
}