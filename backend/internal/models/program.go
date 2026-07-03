package models

import "time"

type Program struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	Code           string       `json:"code" gorm:"size:50"`
	Name           string       `json:"name" gorm:"size:200"`
	OrganizationID uint         `json:"organization_id" gorm:"not null"`
	Organization   Organization `json:"organization" gorm:"foreignKey:OrganizationID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}
