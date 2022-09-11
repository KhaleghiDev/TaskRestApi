package model

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	Title     string
	Disc      string
	Parant_id uint
	Actor     User `gorm:"foreignKey:ID"`
	Status    uint8
}
