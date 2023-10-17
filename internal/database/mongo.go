package database

import (
	"context"
	"fmt"

	"github.com/rparmer/remote-falcon-api/internal/config"
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

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbConfig.Name)

	return &MongoDB{
		Database: db,
	}, nil
}

func (m *MongoDB) Save(collection string, document interface{}) error {
	_, err := m.Database.Collection(collection).InsertOne(context.Background(), document)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) Get(collection string, filter interface{}, result interface{}) error {
	err := m.Database.Collection(collection).FindOne(context.Background(), filter).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) List(collection string, filter interface{}, result interface{}) error {
	cursor, err := m.Database.Collection(collection).Find(context.Background(), filter)
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(context.Background(), result); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) Delete(collection string, filter interface{}) error {
	_, err := m.Database.Collection(collection).DeleteMany(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
