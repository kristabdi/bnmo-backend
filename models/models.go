package models

import (
	"github.com/google/uuid"
	"time"
)

type Default struct {
	ID        uint      `json:"id,omitempty" faker:"-" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at,omitempty" faker:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty" faker:"-"`
}

type User struct {
	Default
	Username   string `json:"username" faker:"username,unique" gorm:"unique"`
	Name       string `json:"name,omitempty" faker:"name,unique"`
	Password   string `json:"password,omitempty" faker:"-"`
	IsVerified bool   `json:"is_verified,omitempty" faker:"-" gorm:"default:false"`
	IsAdmin    bool   `json:"is_admin,omitempty" faker:"-" gorm:"default:false"`
	Balance    uint64 `json:"balance,omitempty" faker:"-" gorm:"default:0"`
	Photo      string `json:"photo,omitempty" faker:"-" gorm:"default:''"`
}

func (u *User) NoSensitive() {
	u.ID = 0
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
	Currency   string    `json:"currency"`
	IsAdd      bool      `json:"is_add"`
	IsApproved bool      `json:"is_approved" gorm:"default:false"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	User       User      `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}
