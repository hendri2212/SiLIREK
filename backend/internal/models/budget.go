package models

import "time"

type Budget struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	UserID        *uint     `json:"user_id" gorm:"uniqueIndex:idx_user_quarter_activity_parent"`
	User          User      `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	QuarterlyID   uint      `json:"quarterly_id" gorm:"not null;uniqueIndex:idx_user_quarter_activity_parent"`
	Quarterly     Quarterly `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ActivityID    uint      `json:"activity_id" gorm:"not null;uniqueIndex:idx_user_quarter_activity_parent"`
	Activity      Activity  `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ParentID      *uint     `json:"parent_id" gorm:"uniqueIndex:idx_user_quarter_activity_parent"`
	Parent        *Budget   `gorm:"foreignKey:ParentID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	SubActivityID *uint     `json:"sub_activity_id" gorm:"-"`
	Budget        float64   `json:"budget"`
	UsedBudget    float64   `json:"used_budget"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (b *Budget) GetRemainingBudget() float64 {
	return b.Budget - b.UsedBudget
}

type BudgetResponse struct {
	ID              uint     `json:"id"`
	UserID          uint     `json:"user_id"`
	QuarterlyID     uint     `json:"quarterly_id"`
	ActivityID      uint     `json:"activity_id"`
	Activity        Activity `json:"activity"`
	ParentID        *uint    `json:"parent_id"`
	Budget          float64  `json:"budget"`
	UsedBudget      float64  `json:"used_budget"`
	RemainingBudget float64  `json:"remaining_budget"`
}
