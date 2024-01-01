package DB

import "go.mongodb.org/mongo-driver/mongo"

type IUserVideosTable interface {
}

func getUserVideosTable(userVideosCollection *mongo.Collection) IUserVideosTable {
	return &userVideosTable{userVideosCollection}
}

type userVideosTable struct {
	table *mongo.Collection
}
