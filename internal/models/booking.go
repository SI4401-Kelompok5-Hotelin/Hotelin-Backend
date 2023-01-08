package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

)

type Booking struct {
		ID			string			`json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
		UserID 	string			`json:"user_id"`
		User		User				`json:"user" gorm:"foreignKey:UserID"`

}

func (u *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}