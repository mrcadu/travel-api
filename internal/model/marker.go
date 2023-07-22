package model

type Marker struct {
	Id         string   `bson:"_id,omitempty" json:"id"`
	Latitude   string   `bson:"latitude,omitempty" json:"latitude"`
	Name       string   `bson:"name,omitempty" json:"name"`
	Longitude  string   `bson:"longitude,omitempty" json:"longitude"`
	MarkerType string   `bson:"type,omitempty" json:"type,omitempty"`
	CountryId  string   `bson:"country_id,omitempty" json:"country_id"`
	Pictures   []string `bson:"pictures,omitempty" json:"pictures"`
}
