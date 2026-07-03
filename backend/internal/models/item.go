package models

import "time"

type Item struct {
	ID                   uint               `json:"id" gorm:"primaryKey"`
	Code                 string             `json:"code" gorm:"size:50"`
	Date                 time.Time          `json:"date" gorm:"type:date"`
	Description          string             `json:"description" gorm:"type:text"`
	Credit               float64            `json:"credit"`
	ExpenditureAccountID uint               `json:"expenditure_account_id" gorm:"not null"`
	ExpenditureAccount   ExpenditureAccount `json:"expenditure_account" gorm:"foreignKey:ExpenditureAccountID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	CreatedAt            time.Time          `json:"created_at"`
	UpdatedAt            time.Time          `json:"updated_at"`
}
