package repositories

import (
	"context"
	"time"

	"github.com/lbAntoine/mongoapi_boilerplate/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository[T models.Model] interface {
	Create(ctx context.Context, model T) error
	FindByID(ctx context.Context, id primitive.ObjectID) (T, error)
	FindOne(ctx context.Context, filter interface{}) (T, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]T, error)
	Update(ctx context.Context, id primitive.ObjectID, update interface{}) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type MongoRepository[T models.Model] struct {
	col *mongo.Collection
}

func NewMongoRepository[T models.Model](col *mongo.Collection) Repository[T] {
	return &MongoRepository[T]{
		col: col,
	}
}

func (r *MongoRepository[T]) Create(ctx context.Context, model T) error {
	now := time.Now()
	model.SetCreatedAt(now)
	model.SetUpdatedAt(now)

	result, err := r.col.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		model.SetID(oid)
	}

	return nil
}

func (r *MongoRepository[T]) FindByID(ctx context.Context, id primitive.ObjectID) (T, error) {
	var model T
	err := r.col.FindOne(ctx, primitive.M{"_id": id}).Decode(&model)
	return model, err
}

func (r *MongoRepository[T]) FindOne(ctx context.Context, filter interface{}) (T, error) {
	var model T
	err := r.col.FindOne(ctx, filter).Decode(&model)
	return model, err
}

func (r *MongoRepository[T]) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]T, error) {
	cur, err := r.col.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var models []T
	if err = cur.All(ctx, &models); err != nil {
		return nil, err
	}

	return models, nil
}

func (r *MongoRepository[T]) Update(ctx context.Context, id primitive.ObjectID, update interface{}) error {
	updateDoc := bson.M{"$set": update}
	updateDoc["$set"].(bson.M)["updated_at"] = time.Now()

	_, err := r.col.UpdateOne(ctx, bson.M{"_id": id}, updateDoc)
	return err
}

func (r *MongoRepository[T]) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.col.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
