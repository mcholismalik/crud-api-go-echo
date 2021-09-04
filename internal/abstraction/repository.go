package abstraction

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type IRepository interface {
	CheckTrx(ctx *Context) *gorm.DB
	Filter(ctx *Context, query *gorm.DB, payload interface{}) *gorm.DB
}

type Repository struct {
	Connection *gorm.DB
	Db         *gorm.DB
	Tx         *gorm.DB
}

func (r *Repository) CheckTrx(ctx *Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}

func (r *Repository) Filter(ctx *Context, query *gorm.DB, payload interface{}) *gorm.DB {
	mVal := reflect.ValueOf(payload)
	mType := reflect.TypeOf(payload)

	for i := 0; i < mVal.NumField(); i++ {
		mValChild := mVal.Field(i)
		mTypeChild := mType.Field(i)

		for j := 0; j < mValChild.NumField(); j++ {
			val := mValChild.Field(j)

			if !val.IsNil() {
				if val.Kind() == reflect.Ptr {
					val = mValChild.Field(j).Elem()
				}

				key := mTypeChild.Type.Field(j).Tag.Get("query")
				filter := mTypeChild.Type.Field(j).Tag.Get("filter")
				// TODO need to custom field
				// filterColumn := mTypeChild.Type.Field(j).Tag.Get("filterColumn")

				switch filter {
				case "LIKE":
					query = query.Where(fmt.Sprintf("%s LIKE ?", key), "%"+val.String()+"%")
				case "ILIKE":
					query = query.Where(fmt.Sprintf("%s ILIKE ?", key), "%"+val.String()+"%")
				case "DATE":
					// TODO we need build custom type first
					// dateStart, dateEnd := date.StringDateToDateRange(val.String())
					// query = query.Where(fmt.Sprintf("%s >= ? and %s <= ?", filterColumn, filterColumn), dateStart, dateEnd)

				default:
					query = query.Where(fmt.Sprintf("%s = ?", key), val.Interface())
				}
			}
		}
	}

	return query
}
