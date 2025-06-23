package models

import "time"

type Quarterly struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Quarter   uint8     `json:"quarter"`
	Year      uint16    `json:"year"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
