package models

type Post struct {
	ID       string `bson:"_id,omitempty"`
	Caption  string `bson:"caption"`
	ImageUrl string `bson:"imageUrl"`
	PostTime string `bson:"time"`
	UID      string `bson:"uid"`
}
