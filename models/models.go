package models

import (
	"github.com/google/uuid"
	"time"
)

type Default struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	Default    `faker:"-"`
	Email      string `json:"email" faker:"email,unique" gorm:"unique"`
	Name       string `json:"name" faker:"name,unique"`
	Password   string `json:"password" faker:"-"`
	IsVerified bool   `json:"is_verified" faker:"-" gorm:"default:false"`
	IsAdmin    bool   `json:"is_admin" faker:"-" gorm:"default:false"`
	Balance    uint64 `json:"balance" faker:"-" gorm:"default:0"`
	Photo      string `json:"photo" faker:"-" gorm:"default:''"`
}

func (u *User) NoPass() {
	u.Password = ""
}

type Transaction struct {
	Default
	IdFrom   int64  `json:"id_from"`
	IdTo     int64  `json:"id_to"`
	Amount   uint64 `json:"amount"`
	UserFrom User   `json:"-" gorm:"foreignKey:IdFrom;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	UserTo   User   `json:"-" gorm:"foreignKey:IdTo;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type Request struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID     int64     `json:"id_user"`
	Amount     uint64    `json:"amount"`
	IsAdd      bool      `json:"isAdd"`
	IsApproved bool      `json:"is_approved" gorm:"default:false"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	User       User      `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}
