package datasource

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"travel-api/internal/config"
)

type MongoDatasource interface {
	SetupMongo()
	CreateMongoIndexes()
	GetClient() *mongo.Client
}

type MongoDatasourceImpl struct {
	client *mongo.Client
}

var client *mongo.Client

func (m MongoDatasourceImpl) GetClient() *mongo.Client {
	return client
}

func (m MongoDatasourceImpl) Setup() {
	m.SetupMongo()
}

func (m MongoDatasourceImpl) SetupMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(config.GetProperty("DB_URL")).SetServerAPIOptions(serverAPI)
	mongoClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	client = mongoClient
	m.CreateMongoIndexes()
}

func (m MongoDatasourceImpl) CreateMongoIndexes() {
	_, err := client.Database("auth").Collection("Profile").Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    bson.D{{"name", 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		panic(err)
	}
	_, err = client.Database("auth").Collection("User").Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    bson.D{{"username", 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		panic(err)
	}
}
