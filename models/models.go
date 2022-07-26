package models

import (
	"time"

	"github.com/google/uuid"
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
	IdFrom       uint   `json:"-"`
	IdTo         uint   `json:"-"`
	Amount       uint64 `json:"amount"`
	UsernameTo   string `json:"username_to" gorm:"-"`
	CurrencyFrom string `json:"currency_from" gorm:"-"`
	NameFrom     string `json:"name_from,omitempty" gorm:"-"`
	NameTo       string `json:"name_to,omitempty" gorm:"-"`
	UserFrom     User   `json:"-" gorm:"foreignKey:IdFrom;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	UserTo       User   `json:"-" gorm:"foreignKey:IdTo;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type Request struct {
	ID         uuid.UUID `json:"id,omitempty" gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID     uint      `json:"id_user,omitempty"`
	Amount     uint64    `json:"amount"`
	Currency   string    `json:"currency" gorm:"-"`
	IsAdd      bool      `json:"is_add"`
	IsApproved bool      `json:"is_approved" gorm:"default:false"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	User       User      `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type QueryType struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

type InfoType struct {
	Timestamp time.Time `json:"timestamp"`
	Rate      float64   `json:"rate"`
}

type Converter struct {
	Success    bool      `json:"success"`
	Query      QueryType `json:"query"`
	Info       InfoType  `json:"info"`
	Historical bool      `json:"historical,omitempty"`
	Date       time.Time `json:"date"`
	Result     float64   `json:"result"`
}
