package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"travel-api/internal/model"
	"travel-api/internal/model/datasource"
)

type Marker interface {
	CRUD[model.Marker]
}

type MarkerMongo struct {
}

func (m MarkerMongo) Get(id string) (model.Marker, error) {
	var marker model.Marker
	err := m.getCollection().FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&marker)
	return marker, err
}

func (m MarkerMongo) Create(marker model.Marker) (model.Marker, error) {
	marker.Id = primitive.NewObjectID().Hex()
	_, err := m.getCollection().InsertOne(context.TODO(), marker)
	return marker, err
}

func (m MarkerMongo) Update(marker model.Marker) (model.Marker, error) {
	updateResult, err := m.getCollection().ReplaceOne(context.TODO(), bson.D{{"_id", marker.Id}}, marker)
	if updateResult != nil && updateResult.ModifiedCount == 0 {
		return marker, mongo.ErrNoDocuments
	}
	return marker, err
}

func (m MarkerMongo) Delete(id string) (string, error) {
	deleteResult, err := m.getCollection().DeleteOne(context.TODO(), bson.D{{"_id", id}})
	if deleteResult != nil && deleteResult.DeletedCount == 0 {
		return id, mongo.ErrNoDocuments
	}
	return id, err
}

func (m MarkerMongo) getCollection() *mongo.Collection {
	return datasource.GetMongoDatasource().GetClient().Database("travel").Collection("Marker")
}

func NewMarkerMongoRepository() MarkerMongo {
	return MarkerMongo{}
}
