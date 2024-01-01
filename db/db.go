package DB

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongo_url              = "mongodb://root:example@localhost:27017/"
	db_name                = "videosCore"
	auth_table_name        = "auth"
	user_videos_table_name = "user-videos"
)

func InitDB() (IAuthTable, IUserVideosTable) {
	client := createConnection()
	db := client.Database(db_name)
	authCollection := db.Collection(auth_table_name)
	userVideosCollection := db.Collection(user_videos_table_name)
	return getAuthTable(authCollection), getUserVideosTable(userVideosCollection)
}

func createConnection() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_url))
	if err != nil {
		fmt.Println("error creating connection :", err)
		panic(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		fmt.Println("error creating connection :", err)
		panic(err)
	}
	return client
}
