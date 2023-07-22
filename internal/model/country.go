package model

type Country struct {
	Id         string `bson:"_id" json:"id"`
	Name       string `bson:"name" json:"name"`
	PictureUrl string `bson:"pictureUrl" json:"pictureUrl"`
}
