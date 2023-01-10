package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	ID      	string 		`json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	UserID		string		`json:"user_id"`
	UserName	string		`json:"user_name"`
	User			User			`json:"user" gorm:"foreignKey:UserID"`
	HotelID		string		`json:"hotel_id"`
	Hotel			Hotel			`json:"hotel" gorm:"foreignKey:HotelID"`
	Review		string		`json:"ulasan"`
}

func (s *Review) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return
}

type ReviewRequest struct {
	HotelID		string	`json:"hotel_id"`
	Review		string	`json:"ulasan"`
}

type ReviewResponse struct {
	ID				string	`json:"id"`
	UserName	string	`json:"user_name"`
	Review		string	`json:"ulasan"`
}