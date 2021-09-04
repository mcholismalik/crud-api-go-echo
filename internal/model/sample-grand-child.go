package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/util/date"

	"gorm.io/gorm"
)

type SampleGrandChildEntity struct {
	Key   string `json:"key" validate:"required" gorm:"index:idx_sample_grand_child_key,unique"`
	Value string `json:"value" validate:"required"`
}

type SampleGrandChildFilter struct {
	Key   *string `query:"key" filter:"ILIKE"`
	Value *string `query:"value"`
}

type SampleGrandChildEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	SampleGrandChildEntity

	// relations
	SampleChildId int `json:"sample_child_id"`

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type SampleGrandChildFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	SampleGrandChildFilter
}

func (SampleGrandChildEntityModel) TableName() string {
	return "sample_grand_childs"
}

func (m *SampleGrandChildEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *SampleGrandChildEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
