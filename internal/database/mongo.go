package database

import (
	"context"
	"fmt"
	"time"

	"github.com/rparmer/remote-falcon-api/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Database *mongo.Database
}

func New() (*MongoDB, error) {
	dbConfig := config.GetConfig().Database
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	credential := options.Credential{
		Username: dbConfig.Username,
		Password: dbConfig.Password,
	}
	uri := fmt.Sprintf("mongodb://%v", dbConfig.Address)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI).SetAuth(credential)
	ctx, cancel := newContextTimeout()
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	db := client.Database(dbConfig.Name)

	for _, collection := range collections {
		if err := db.CreateCollection(ctx, collection.Name, collection.CreateCollectionOptions); err != nil {
			return nil, err
		}
		if collection.Indexes != nil {
			for i, index := range *collection.Indexes {
				indexModel := mongo.IndexModel{
					Keys:    bson.D{{Key: index, Value: i + 1}},
					Options: options.Index().SetUnique(true),
				}
				db.Collection(collection.Name).Indexes().CreateOne(ctx, indexModel)
			}
		}
	}

	return &MongoDB{
		Database: db,
	}, nil
}

func (m *MongoDB) Create(collection string, document interface{}) error {
	ctx, cancel := newContextTimeout()
	defer cancel()
	_, err := m.Database.Collection(collection).InsertOne(ctx, document)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) Update(collection string, filter interface{}, document interface{}) (int64, error) {
	update := bson.D{{Key: "$set", Value: document}}
	ctx, cancel := newContextTimeout()
	defer cancel()
	res, err := m.Database.Collection(collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err

	}
	return res.MatchedCount, nil
}

func (m *MongoDB) BulkWrite(collection string, writeModel []mongo.WriteModel) error {
	ctx, cancel := newContextTimeout()
	defer cancel()
	_, err := m.Database.Collection(collection).BulkWrite(ctx, writeModel)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) Get(collection string, filter interface{}, result interface{}) error {
	ctx, cancel := newContextTimeout()
	defer cancel()
	err := m.Database.Collection(collection).FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) List(collection string, filter interface{}, result interface{}) error {
	ctx, cancel := newContextTimeout()
	defer cancel()
	cursor, err := m.Database.Collection(collection).Find(ctx, filter)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, result); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) Delete(collection string, filter interface{}) error {
	ctx, cancel := newContextTimeout()
	defer cancel()
	_, err := m.Database.Collection(collection).DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func newContextTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*5)
}
