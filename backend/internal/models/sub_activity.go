package models

import "time"

type SubActivity struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Code       string    `json:"code" gorm:"size:50"`
	Name       string    `json:"name" gorm:"size:200"`
	ActivityID uint      `json:"activity_id" gorm:"not null"`
	Activity   Activity  `json:"activity" gorm:"foreignKey:ActivityID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
