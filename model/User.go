package model

import (
	"time"
)


type User struct {
	ID      uint `gorm:"primarykey"`
	Username string
	Email    string
	Password string
	Role     uint8
	Status   uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Users []User
