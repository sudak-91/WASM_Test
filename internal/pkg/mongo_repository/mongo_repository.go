package mongorepository

import (
	"github.com/sudak-91/wasm-test/pkg/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	Users repository.Users
}

func NewMongoRepository(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		Users: NewMongoUser(db),
	}
}
