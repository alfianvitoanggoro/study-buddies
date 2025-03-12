package factory

import (
	"context"

	"github.com/AlfianVitoAnggoro/study-buddies/internal/repository"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

type Factory struct {
	Db              *gorm.DB
	RDb             *redis.Client
	Ctx             context.Context
	Es              *elasticsearch.Client
	RedisRepository repository.RedisRepository
	UserRepository  repository.User
}

func NewFactory(db *gorm.DB, ctx context.Context, rdb *redis.Client, es *elasticsearch.Client) *Factory {
	f := &Factory{}
	f.SetupDb(db)
	f.SetupRedisDb(ctx, rdb)
	f.SetupRedisDb(ctx, rdb)
	f.SetupRepository()

	return f
}

func (f *Factory) SetupDb(db *gorm.DB) {
	f.Db = db

}
func (f *Factory) SetupRedisDb(ctx context.Context, rdb *redis.Client) {
	f.Ctx = ctx
	f.RDb = rdb
}

func (f *Factory) SetupElasticsearch(es *elasticsearch.Client) {
	f.Es = es
}

func (f *Factory) SetupRepository() {
	if f.Db == nil {
		panic("Failed setup repository, db is undefined")
	}

	f.RedisRepository = repository.NewRedis(f.Ctx, f.RDb)
	f.UserRepository = repository.NewUser(f.Db)
}
