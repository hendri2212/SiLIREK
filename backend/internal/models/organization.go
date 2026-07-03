package models

import "time"

type Organization struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"type:varchar(200);not null"`
	Number    string         `json:"number" gorm:"type:varchar(50);not null"`
	ParentID  *uint          `json:"parent_id" gorm:"index;default:NULL"`
	Children  []Organization `json:"children" gorm:"foreignKey:ParentID"`
	Parent    *Organization  `json:"parent" gorm:"foreignKey:ParentID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
