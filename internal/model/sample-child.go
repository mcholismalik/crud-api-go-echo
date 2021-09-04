package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/util/date"

	"gorm.io/gorm"
)

type SampleChildEntity struct {
	Key   string `json:"key" validate:"required" gorm:"index:idx_sample_child_key,unique"`
	Value string `json:"value" validate:"required"`
}

type SampleChildFilter struct {
	Key   *string `query:"key" filter:"ILIKE"`
	Value *string `query:"value"`
}

type SampleChildEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	SampleChildEntity

	// relations
	SampleId          int                           `json:"sample_id"`
	SampleGrandChilds []SampleGrandChildEntityModel `json:"sample_grand_childs" gorm:"foreignKey:SampleChildId"`

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type SampleChildFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	SampleChildFilter
}

func (SampleChildEntityModel) TableName() string {
	return "sample_childs"
}

func (m *SampleChildEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *SampleChildEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
