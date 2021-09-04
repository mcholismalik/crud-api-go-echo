package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/model"
	res "codeid-boiler/pkg/util/response"
)

// Get
type SampleGetRequest struct {
	abstraction.Pagination
	model.SampleFilterModel
}
type SampleGetResponse struct {
	Datas          []model.SampleEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type SampleGetResponseDoc struct {
	Body struct {
		Meta res.Meta                  `json:"meta"`
		Data []model.SampleEntityModel `json:"data"`
	} `json:"body"`
}

// GetByID
type SampleGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type SampleGetByIDResponse struct {
	model.SampleEntityModel
}
type SampleGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data SampleGetByIDResponse `json:"data"`
	} `json:"body"`
}

// Create
type SampleCreateRequest struct {
	model.SampleEntity
}
type SampleCreateResponse struct {
	model.SampleEntityModel
}
type SampleCreateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data SampleCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type SampleUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.SampleEntity
}
type SampleUpdateResponse struct {
	model.SampleEntityModel
}
type SampleUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data SampleUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type SampleDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type SampleDeleteResponse struct {
	model.SampleEntityModel
}
type SampleDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data SampleDeleteResponse `json:"data"`
	} `json:"body"`
}
