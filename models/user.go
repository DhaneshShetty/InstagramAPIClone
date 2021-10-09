package models

type User struct {
	OID      string `bson:"_id,omitempty"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
