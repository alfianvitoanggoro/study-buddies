package factory

import (
	"github.com/AlfianVitoAnggoro/study-buddies/internal/repository"

	"gorm.io/gorm"
)

type Factory struct {
	Db             *gorm.DB
	UserRepository repository.User
}

func NewFactory() *Factory {
	f := &Factory{}
	f.SetupDb()
	f.SetupRepository()

	return f
}

func (f *Factory) SetupDb(Db *gorm.DB) {
	f.Db = Db
}

func (f *Factory) SetupRepository() {
	if f.Db == nil {
		panic("Failed setup repository, db is undefined")
	}

	f.UserRepository = repository.NewUser(f.Db)
}
