package models

import "time"

type Activity struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Code      string `json:"code" gorm:"size:50;unique"`
	Name      string `json:"name" gorm:"size:200;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
