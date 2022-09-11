package model

import "time"

type Response struct {
	ID        uint `gorm:"primarykey"`
	Disc      string
	Ticket    uint
	User_id   User `gorm:"foreignKey:ID"`
	Status    uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}
