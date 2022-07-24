package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" gorm:"unique" faker:"email,unique"`
	Name     string `json:"name" faker:"name,unique"`
	Password string `json:"password" faker:"password"`
	Balance  int64  `json:"balance" gorm:"default:0"`
	IsAdmin  bool   `json:"isAdmin"`
}

type Transaction struct {
	gorm.Model
	IdFrom   int64 `json:"id_from"`
	IdTo     int64 `json:"id_to"`
	UserFrom User  `json:"-" gorm:"foreignKey:IdFrom;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	UserTo   User  `json:"-" gorm:"foreignKey:IdTo;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type Request struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID     int64     `json:"id_user"`
	Amount     uint64    `json:"amount"`
	IsAdd      bool      `json:"isAdd"`
	IsApproved bool      `json:"is_approved" gorm:"default:false"`
	User       User      `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}
