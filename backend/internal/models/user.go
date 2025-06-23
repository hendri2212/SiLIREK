package models

import "time"

// UserRole defines valid roles for a user.
type UserRole string

const (
	UserRoleSuperadmin UserRole = "superadmin"
	UserRoleAdmin      UserRole = "admin"
	UserRoleUser       UserRole = "user"
)

type User struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	FullName   string    `json:"full_name" gorm:"size:50"`
	Email      string    `json:"email" gorm:"unique;size:50"`
	Password   string    `json:"password" gorm:"size:60"`
	Nip        *string   `json:"nip" gorm:"size:25;default:NULL"`
	PositionID *uint     `json:"position_id" gorm:"default:NULL"`
	Position   *Position `json:"position" gorm:"foreignKey:PositionID;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	LeaderID   *uint     `json:"leader_id" gorm:"default:NULL"`
	Leader     *Leader   `json:"leader" gorm:"foreignKey:LeaderID;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Role       UserRole  `json:"role" gorm:"type:ENUM('superadmin','admin','user');default:'user'"`
	Photo      *string   `json:"photo" gorm:"default:NULL"`
	IsActive   bool      `json:"is_active" gorm:"default:true"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
