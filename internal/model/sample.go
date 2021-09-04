package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/util/date"

	"gorm.io/gorm"
)

type SampleEntity struct {
	Key   string `json:"key" validate:"required" gorm:"index:idx_sample_key,unique"`
	Value string `json:"value" validate:"required"`
}

type SampleFilter struct {
	Key   *string `query:"key" filter:"ILIKE"`
	Value *string `query:"value"`
}

type SampleEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	SampleEntity

	// relations
	SampleChilds []SampleChildEntityModel `json:"sample_childs" gorm:"foreignKey:SampleId"`

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type SampleFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	SampleFilter
}

func (SampleEntityModel) TableName() string {
	return "samples"
}

func (m *SampleEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *SampleEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
