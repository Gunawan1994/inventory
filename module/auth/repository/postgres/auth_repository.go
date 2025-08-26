package postgres

import (
	"context"
	"errors"
	"time"

	"inventory-service/domain"
	"inventory-service/helpers/utils"
	"inventory-service/model"

	"gorm.io/gorm"
)

type AuthRepo struct {
}

func NewAuthRepository() AuthIRepository {
	return &AuthRepo{}
}

func (auth *AuthRepo) VerifyCredential(ctx context.Context, tx *gorm.DB, req model.VerifyCredential) (*domain.User, error) {
	var user *domain.User
	getUser := tx.Find(&user, "email = ?", req.Email)

	if getUser.Error != nil {
		return nil, getUser.Error
	}

	if utils.CompareToHash(req.Password, user.Password) {
		exp := time.Now().Add(time.Hour * 24)
		token, err := utils.GenerateToken(&model.VerifyCredentialRes{
			User: *user,
		}, "user", exp.Unix())
		if err != nil {
			return nil, err
		}

		user.Token = token

		return user, nil
	} else {
		return nil, errors.New("unmatch credential, please check password")
	}
}
