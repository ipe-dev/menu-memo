package model

import "time"

type User struct {
	Id uint `gorm:"primary_key"`
	Name string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}