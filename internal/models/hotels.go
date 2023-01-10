package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Hotel struct {
	ID      	string 		`json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name    	string 		`json:"name"`
	Email   	string 		`json:"email"`
	Phone   	string 		`json:"phone"`
	Address 	string 		`json:"address"`
	Rating		string		`json:"rating"`
	ListRoom 	[]Room		`json:"list_room" gorm:"foreignKey:RoomID"`	
}

func (s *Hotel) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return
}

type HotelRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Rating	string `json:"rating"`
}

type HotelResponse struct {
	ID      	string `json:"id"`
	Name    	string `json:"name"`
	Email   	string `json:"email"`
	Phone   	string `json:"phone"`
	Address 	string `json:"address"`
	Rating		string `json:"rating"`
	ListRoom	
}