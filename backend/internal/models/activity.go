package models

import "time"

type Activity struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Code          string        `json:"code" gorm:"size:50"`
	Name          string        `json:"name" gorm:"size:200"`
	ProgramID     uint          `json:"program_id" gorm:"not null"`
	Program       Program       `json:"program" gorm:"foreignKey:ProgramID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	SubActivities []SubActivity `json:"sub_activities,omitempty" gorm:"foreignKey:ActivityID"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}
