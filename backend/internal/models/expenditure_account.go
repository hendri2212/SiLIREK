package models

import "time"

type ExpenditureAccount struct {
	ID            uint        `json:"id" gorm:"primaryKey"`
	Code          string      `json:"code" gorm:"size:50"`
	Description   string      `json:"description" gorm:"type:text"`
	BudgetCeiling float64     `json:"budget_ceiling"`
	SubActivityID uint        `json:"sub_activity_id" gorm:"not null"`
	SubActivity   SubActivity `json:"sub_activity" gorm:"foreignKey:SubActivityID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}
