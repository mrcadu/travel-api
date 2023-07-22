package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"travel-api/internal/model"
	"travel-api/internal/model/datasource"
)

type Country interface {
	CRUD[model.Country]
	GetMarkers(id string) ([]model.Marker, error)
	GetAll() ([]model.Country, error)
}

type CountryMongo struct {
}

func (c CountryMongo) Get(id string) (model.Country, error) {
	var country model.Country
	err := c.getCollection().FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&country)
	return country, err
}

func (c CountryMongo) GetAll() ([]model.Country, error) {
	var countries []model.Country
	result, err := c.getCollection().Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, nil
	}
	err = result.All(context.TODO(), &countries)
	if err != nil {
		return nil, nil
	}
	return countries, err
}

func (c CountryMongo) Create(country model.Country) (model.Country, error) {
	country.Id = primitive.NewObjectID().Hex()
	_, err := c.getCollection().InsertOne(context.TODO(), country)
	return country, err
}

func (c CountryMongo) Update(country model.Country) (model.Country, error) {
	updateResult, err := c.getCollection().ReplaceOne(context.TODO(), bson.D{{"_id", country.Id}}, country)
	if updateResult != nil && updateResult.ModifiedCount == 0 {
		return country, mongo.ErrNoDocuments
	}
	return country, err
}

func (c CountryMongo) Delete(id string) (string, error) {
	deleteResult, err := c.getCollection().DeleteOne(context.TODO(), bson.D{{"_id", id}})
	if deleteResult != nil && deleteResult.DeletedCount == 0 {
		return id, mongo.ErrNoDocuments
	}
	return id, err
}

func (c CountryMongo) GetMarkers(id string) ([]model.Marker, error) {
	var markers []model.Marker
	result, err := c.getMarkerCollection().Find(context.TODO(), bson.D{{"country_id", id}})
	if err != nil {
		return nil, nil
	}
	err = result.All(context.TODO(), &markers)
	if err != nil {
		return nil, nil
	}
	return markers, err
}

func (c CountryMongo) getCollection() *mongo.Collection {
	return datasource.GetMongoDatasource().GetClient().Database("travel").Collection("Country")
}

func (c CountryMongo) getMarkerCollection() *mongo.Collection {
	return datasource.GetMongoDatasource().GetClient().Database("travel").Collection("Marker")
}

func NewCountryMongoRepository() CountryMongo {
	return CountryMongo{}
}
