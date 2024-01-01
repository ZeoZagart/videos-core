package DB

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Email        string `bson:"_id,omitempty" json:"email"`
	Username     string `bson:"username,omitempty"`
	PasswordHash string `bson:"password,omitempty"`
}

type UserVideos struct {
	Email    string               `bson:"_id,omitempty" json:"email"`
	VideoIDs []primitive.ObjectID `bson:"videoids,omitempty"`
}

type Video struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}
