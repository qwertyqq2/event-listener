package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/qwertyqq2/event-listener/pkg/mongodb/event"
)

var (
	collection *mongo.Collection
)

func init() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://mainuser:111@eventsclaster.bvnv49y.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("tasker").Collection("tasks")
	log.Println("Mongo is connect")
}

func Insert(ctx context.Context, e event.Event) error {
	_, err := collection.InsertOne(ctx, e)
	return err
}

func GetAll(ctx context.Context) error {
	return nil
}
