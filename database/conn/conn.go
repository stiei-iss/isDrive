package conn

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(url string) {
	fmt.Println("--")
	clientOpt := options.Client().ApplyURI(url)
	client, _ := mongo.Connect(nil, clientOpt)

	if err := client.Ping(nil, nil); err != nil {
		panic(err)
	}
	fmt.Println("conn succeed!")
}
