package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	collection *mongo.Collection
)

type Event struct {
	Index int `bson:"index"`
	Value int `bson:"value"`
}

func CreateEvent(ctx context.Context, e *Event) error {
	_, err := collection.InsertOne(ctx, e)
	return err
}

func main() {

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://mainuser:111@eventsclaster.bvnv49y.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("tasker").Collection("tasks")

	e1 := &Event{
		Index: 2,
		Value: 2,
	}

	e2 := &Event{
		Index: 3,
		Value: 3,
	}

	e3 := &Event{
		Index: 4,
		Value: 4,
	}

	err = CreateEvent(ctx, e1)
	if err != nil {
		log.Fatal(err)
	}
	err = CreateEvent(ctx, e2)
	if err != nil {
		log.Fatal(err)
	}
	err = CreateEvent(ctx, e3)
	if err != nil {
		log.Fatal(err)
	}
}
