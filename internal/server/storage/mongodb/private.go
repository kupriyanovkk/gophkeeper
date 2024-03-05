package mongodb

import (
	"context"

	"github.com/kupriyanovkk/gophkeeper/internal/server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PrivateMongodbStorage
type PrivateMongodbStorage struct {
	client *mongo.Client
}

// CreatePrivateData handles the private data creation
func (s *PrivateMongodbStorage) CreatePrivateData(ctx context.Context, private model.PrivateData) (model.PrivateData, error) {
	privateCollection := s.client.Database("gophkeeper").Collection("private")

	privateData := bson.D{
		{Key: "user_id", Value: private.UserID},
		{Key: "title", Value: private.Title},
		{Key: "type", Value: private.Type},
		{Key: "content", Value: private.Content},
		{Key: "updated", Value: private.Updated},
		{Key: "deleted", Value: private.Deleted},
	}
	result, err := privateCollection.InsertOne(ctx, privateData)
	if err != nil {
		return model.PrivateData{}, err
	}
	private.ID = result.InsertedID.(uint32)

	return private, nil
}

// UpdatePrivateData handles the private data update
func (s *PrivateMongodbStorage) UpdatePrivateData(ctx context.Context, private model.PrivateData) (model.PrivateData, error) {
	privateCollection := s.client.Database("gophkeeper").Collection("private")

	_, err := privateCollection.UpdateOne(ctx, bson.D{
		{Key: "user_id", Value: private.UserID},
		{Key: "title", Value: private.Title},
		{Key: "type", Value: private.Type},
		{Key: "content", Value: private.Content},
		{Key: "updated", Value: private.Updated},
		{Key: "deleted", Value: private.Deleted},
	}, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "content", Value: private.Content},
			{Key: "updated", Value: private.Updated},
			{Key: "deleted", Value: private.Deleted},
		}},
	})
	if err != nil {
		return model.PrivateData{}, err
	}

	return private, nil
}

func (s *PrivateMongodbStorage) GetPrivateData(ctx context.Context, private model.PrivateData) (model.PrivateData, error) {
	privateCollection := s.client.Database("gophkeeper").Collection("private")

	result := model.PrivateData{}
	err := privateCollection.FindOne(ctx, bson.D{
		{Key: "user_id", Value: private.UserID},
		{Key: "title", Value: private.Title},
		{Key: "type", Value: private.Type},
		{Key: "content", Value: private.Content},
		{Key: "updated", Value: private.Updated},
		{Key: "deleted", Value: private.Deleted},
	}).Decode(&result)
	if err != nil {
		return model.PrivateData{}, err
	}

	return result, nil
}

func (s *PrivateMongodbStorage) DeletePrivateData(ctx context.Context, private model.PrivateData) error {
	privateCollection := s.client.Database("gophkeeper").Collection("private")

	_, err := privateCollection.DeleteOne(ctx, bson.D{
		{Key: "user_id", Value: private.UserID},
		{Key: "title", Value: private.Title},
		{Key: "type", Value: private.Type},
		{Key: "content", Value: private.Content},
		{Key: "updated", Value: private.Updated},
		{Key: "deleted", Value: private.Deleted},
	})

	return err
}

func (s *PrivateMongodbStorage) GetPrivateDataByType(ctx context.Context, privateType model.PrivateDataType, user model.User) ([]model.PrivateData, error) {
	privateCollection := s.client.Database("gophkeeper").Collection("private")

	privateData := []model.PrivateData{}
	cursor, err := privateCollection.Find(ctx, bson.D{{Key: "user_id", Value: user.ID}, {Key: "type", Value: privateType.ID}})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &privateData); err != nil {
		return nil, err
	}

	return privateData, nil
}

// NewPrivateStore creates a new PrivateMongodbStorage instance.
//
// c *mongo.Client
// *PrivateMongodbStorage
func NewPrivateStore(c *mongo.Client) *PrivateMongodbStorage {
	return &PrivateMongodbStorage{
		client: c,
	}
}
