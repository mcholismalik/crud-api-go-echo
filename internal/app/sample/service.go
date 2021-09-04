package sample

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/dto"
	"codeid-boiler/internal/factory"
	"codeid-boiler/internal/model"
	"codeid-boiler/internal/repository"
	res "codeid-boiler/pkg/util/response"
	"codeid-boiler/pkg/util/trxmanager"
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	Find(ctx *abstraction.Context, payload *dto.SampleGetRequest) (*dto.SampleGetResponse, error)
	FindByID(ctx *abstraction.Context, payload *dto.SampleGetByIDRequest) (*dto.SampleGetByIDResponse, error)
	Create(ctx *abstraction.Context, payload *dto.SampleCreateRequest) (*dto.SampleCreateResponse, error)
	Update(ctx *abstraction.Context, payload *dto.SampleUpdateRequest) (*dto.SampleUpdateResponse, error)
	Delete(ctx *abstraction.Context, payload *dto.SampleDeleteRequest) (*dto.SampleDeleteResponse, error)
}

type service struct {
	Repository repository.Sample
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.SampleRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.SampleGetRequest) (*dto.SampleGetResponse, error) {
	var result *dto.SampleGetResponse
	var datas *[]model.SampleEntityModel

	datas, info, err := s.Repository.Find(ctx, &payload.SampleFilterModel, &payload.Pagination)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.SampleGetResponse{
		Datas:          *datas,
		PaginationInfo: *info,
	}

	return result, nil
}

func (s *service) FindByID(ctx *abstraction.Context, payload *dto.SampleGetByIDRequest) (*dto.SampleGetByIDResponse, error) {
	var result *dto.SampleGetByIDResponse

	data, err := s.Repository.FindByID(ctx, &payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.SampleGetByIDResponse{
		SampleEntityModel: *data,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.SampleCreateRequest) (*dto.SampleCreateResponse, error) {
	var result *dto.SampleCreateResponse
	var data *model.SampleEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data.Context = ctx
		data.SampleEntity = payload.SampleEntity
		data, err = s.Repository.Create(ctx, data)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.SampleCreateResponse{
		SampleEntityModel: *data,
	}

	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.SampleUpdateRequest) (*dto.SampleUpdateResponse, error) {
	var result *dto.SampleUpdateResponse
	var data *model.SampleEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		_, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}

		data.Context = ctx
		data.SampleEntity = payload.SampleEntity
		data, err = s.Repository.Update(ctx, &payload.ID, data)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.SampleUpdateResponse{
		SampleEntityModel: *data,
	}

	return result, nil
}

func (s *service) Delete(ctx *abstraction.Context, payload *dto.SampleDeleteRequest) (*dto.SampleDeleteResponse, error) {
	var result *dto.SampleDeleteResponse
	var data *model.SampleEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}

		data.Context = ctx
		data, err = s.Repository.Delete(ctx, &payload.ID, data)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.SampleDeleteResponse{
		SampleEntityModel: *data,
	}

	return result, nil
}
