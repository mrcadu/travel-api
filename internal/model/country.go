package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Country struct {
	Id          string             `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	PictureUrl  string             `bson:"pictureUrl" json:"pictureUrl"`
	Checklist   []string           `bson:"checklist" json:"checklist"`
	InitialDate primitive.DateTime `bson:"initialDate" json:"initialDate"`
	FinalDate   primitive.DateTime `bson:"finalDate" json:"finalDate"`
}
