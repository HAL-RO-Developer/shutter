package model

import (
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type Account struct {
	TwitterID      string
	CustomerKey    string
	CustomerSecret string
	AccessToken    string
	AccessSecret   string
	Alive          bool
	Model
}
