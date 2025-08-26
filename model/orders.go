package model

import (
	"time"

	"inventory-service/domain"
)

type BaseOrders struct {
	Id        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductId int64     `gorm:"not null" json:"product_id"`
	UserId    int64     `gorm:"not null" json:"user_id"`
	Status    string    `gorm:"type:varchar(20);default:PENDING;not null" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy *int64    `json:"created_by,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy *int64    `json:"updated_by,omitempty"`
}

func (req BaseOrders) ToDomain() *domain.Orders {
	return &domain.Orders{
		Id:        req.Id,
		ProductId: req.ProductId,
		UserId:    req.UserId,
		Status:    req.Status,
		CreatedAt: time.Now(),
	}
}

type CreateOrdersReq struct {
	*BaseOrders
}

type CreateOrdersRes struct {
	*domain.Orders
}

type UpdateOrdersReq struct {
	Id int64 `json:"id"`
	*BaseOrders
}

type UpdateOrdersRes struct {
	*domain.Orders
}

type GetListOrdersReq struct {
	Page    PaginationParam
	Filter  FilterParams
	Order   OrderParam
	Keyword KeywordParam
}

type GetListOrdersRes struct {
	Data       []*domain.Orders
	Pagination *Pagination
}

type GetIdOrdersReq struct {
	Id int64 `json:"id"`
}

type GetIdOrdersRes struct {
	*domain.Orders
}

type DeleteOrdersReq struct {
	Id int64 `json:"id"`
}

type DeleteOrdersRes struct {
	*domain.Orders
}
