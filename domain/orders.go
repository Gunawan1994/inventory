package domain

import "time"

type Orders struct {
	Id        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductId int64     `gorm:"not null" json:"product_id"`
	UserId    int64     `gorm:"not null" json:"user_id"`
	Status    string    `gorm:"type:varchar(20);default:PENDING;not null" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy *int64    `json:"created_by,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy *int64    `json:"updated_by,omitempty"`
}
