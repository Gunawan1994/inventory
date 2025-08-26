package domain

import "time"

type Products struct {
	Id        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Quantity  int       `gorm:"not null;default:0" json:"quantity"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy *int64    `json:"created_by,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy *int64    `json:"updated_by,omitempty"`

	// Associations
	Orders []Orders `gorm:"foreignKey:ProductId" json:"orders,omitempty"`
}
