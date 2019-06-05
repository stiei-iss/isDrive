package conn

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestConnectDB(t *testing.T) {
	clientOpt := options.Client().ApplyURI("mongodb://10.0.5.240:27017")
	client, _ := mongo.Connect(nil, clientOpt)

	if err := client.Ping(nil, nil); err != nil {
		t.Log(err)
	}
}
