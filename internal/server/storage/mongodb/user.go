package mongodb

import (
	"context"

	"github.com/kupriyanovkk/gophkeeper/internal/server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongodbStorage struct {
	client *mongo.Client
}

// Create creates a new user in the UserMongodbStorage.
//
// ctx context.Context, user model.User
// model.User, error
func (s *UserMongodbStorage) Create(ctx context.Context, user model.User) (model.User, error) {
	usersCollection := s.client.Database("gophkeeper").Collection("user")
	userData := bson.D{
		{Key: "login", Value: user.Login},
		{Key: "password", Value: user.Password},
	}
	result, err := usersCollection.InsertOne(ctx, userData)
	if err != nil {
		return model.User{}, err
	}
	user.ID = result.InsertedID.(uint32)

	return user, nil
}

// Update updates the user in the MongoDB storage.
//
// ctx context.Context, user model.User
// model.User, error
func (s *UserMongodbStorage) Update(ctx context.Context, user model.User) (model.User, error) {
	usersCollection := s.client.Database("gophkeeper").Collection("user")
	_, err := usersCollection.UpdateOne(ctx, bson.D{
		{Key: "login", Value: user.Login},
		{Key: "password", Value: user.Password},
	}, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "login", Value: user.Login},
			{Key: "password", Value: user.Password},
		}},
	})
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// Get description of the Go function.
//
// ctx context.Context, user model.User.
// model.User, error.
func (s *UserMongodbStorage) Get(ctx context.Context, user model.User) (model.User, error) {
	usersCollection := s.client.Database("gophkeeper").Collection("user")

	result := model.User{}
	err := usersCollection.FindOne(ctx, bson.D{
		{Key: "login", Value: user.Login},
		{Key: "password", Value: user.Password},
	}).Decode(&result)
	if err != nil {
		return model.User{}, err
	}

	return result, nil
}

// NewUserStore initializes a new UserMongodbStorage with the given mongo client.
//
// c *mongo.Client - the mongo client to be used by the UserMongodbStorage.
// *UserMongodbStorage - returns a pointer to the newly initialized UserMongodbStorage.
func NewUserStore(c *mongo.Client) *UserMongodbStorage {
	return &UserMongodbStorage{
		client: c,
	}
}
