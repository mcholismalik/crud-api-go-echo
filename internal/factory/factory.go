package factory

import (
	"codeid-boiler/database"
	"codeid-boiler/internal/repository"

	"gorm.io/gorm"
)

type Factory struct {
	Db               *gorm.DB
	UserRepository   repository.User
	SampleRepository repository.Sample
}

func NewFactory() *Factory {
	f := &Factory{}
	f.SetupDb()
	f.SetupRepository()

	return f
}

func (f *Factory) SetupDb() {
	db, err := database.Connection("SAMPLE1")
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	f.Db = db
}

func (f *Factory) SetupRepository() {
	if f.Db == nil {
		panic("Failed setup repository, db is undefined")
	}

	f.UserRepository = repository.NewUser(f.Db)
	f.SampleRepository = repository.NewSample(f.Db)
}
