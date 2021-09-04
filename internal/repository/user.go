package repository

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/model"

	"gorm.io/gorm"
)

type User interface {
	FindByEmail(ctx *abstraction.Context, email *string) (*model.UserEntityModel, error)
	Create(ctx *abstraction.Context, data *model.UserEntity) (*model.UserEntityModel, error)
	checkTrx(ctx *abstraction.Context) *gorm.DB
}

type user struct {
	abstraction.Repository
}

func NewUser(db *gorm.DB) *user {
	return &user{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *user) FindByEmail(ctx *abstraction.Context, email *string) (*model.UserEntityModel, error) {
	conn := r.checkTrx(ctx)

	var data model.UserEntityModel
	err := conn.Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *user) Create(ctx *abstraction.Context, e *model.UserEntity) (*model.UserEntityModel, error) {
	conn := r.checkTrx(ctx)

	var data model.UserEntityModel
	data.UserEntity = *e
	err := conn.Create(&data).Error
	if err != nil {
		return nil, err
	}
	err = conn.Model(&data).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *user) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
