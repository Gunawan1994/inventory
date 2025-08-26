package model

import (
	"inventory-service/domain"
	"time"
)

type BaseUser struct {
	Id        int64     `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	Id        int64     `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (req *BaseUser) ToDomain() *domain.User {
	return &domain.User{
		Id:       req.Id,
		Email:    req.Email,
		Password: req.Password,
		Username: req.Username,
	}
}

type CreateUserReq struct {
	*BaseUser
}

type CreateUserRes struct {
	*domain.User
}
