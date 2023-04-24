package eventstores

import (
	"context"
	"fmt"
	"time"

	"github.com/dritelabs/blog-reactive/internal/blog/infrastructure/event_stores/models"
	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoEventStore struct {
	eventCollection *mongo.Collection
}

func (es *MongoEventStore) List(ctx context.Context) ([]models.Event, error) {
	var events []models.Event

	cur, err := es.eventCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", cur.Next(ctx))

	if err = cur.All(ctx, &events); err != nil {
		panic(err)
	}

	for _, event := range events {
		err = cur.Decode(&event)
		if err != nil {
			panic(err)
		}
	}

	return events, nil
}

func (es *MongoEventStore) Store(ctx context.Context, e domain.Event) (*models.Event, error) {
	_, err := es.eventCollection.InsertOne(ctx, models.Event{
		CreatedAt: time.Now(),
		Payload:   e,
		Type:      e.String(),
	})
	if err != nil {
		return nil, err
	}

	fmt.Print(es.List(ctx))

	return nil, nil
}

func NewMongoEventStore(client *mongo.Client) *MongoEventStore {
	eventCollection := client.Database("events").Collection("streams")

	return &MongoEventStore{
		eventCollection: eventCollection,
	}
}
