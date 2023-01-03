package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserType struct {
	ID     string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	UserID string    `json:"user_id"`
	User   User 			`json:"user" gorm:"foreignKey:UserID"`
	TypeID string    `json:"type_id"`
	Type   Type      `json:"type" gorm:"foreignKey:TypeID"`
}

func (u *UserType) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}
