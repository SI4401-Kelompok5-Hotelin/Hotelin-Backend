package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserBalance struct {
	ID      string  `json:"id" gorm:"primary_key, type:uid, default:uuid_generate_v4()"`
	UserID  string  `json:"user_id"`
	User    User    `json:"user" gorm:"foreignKey:UserID"`
	Balance float64 `json:"balance" default:"0"`
}

func (u *UserBalance) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}

type UserBalanceRequest struct {
	Balance float64 `json:"balance" validate:"required"`
}