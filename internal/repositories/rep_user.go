package repositories

import (
	"context"

	"github.com/lbAntoine/mongoapi_boilerplate/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Repository[*models.User]
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	UpdatePassword(ctx context.Context, id string, password string) error
}

type MongoUserRepository struct {
	*MongoRepository[*models.User]
}

func NewUserRepository(col *mongo.Collection) UserRepository {
	return &MongoUserRepository{
		MongoRepository: NewMongoRepository[*models.User](col).(*MongoRepository[*models.User]),
	}
}

func (r *MongoUserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	return r.FindOne(ctx, bson.M{"email": email})
}

func (r *MongoUserRepository) UpdatePassword(ctx context.Context, id string, password string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{"password": password}
	return r.Update(ctx, oid, update)
}
