package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDatabase = "gif_engine" // @todo parameterise
var MongoCollection = "images"

func NewMongoDB(ctx context.Context, connectionURI string) (DB, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		return nil, err
	}

	coll := client.Database(MongoDatabase).Collection(MongoCollection)

	return &MongoDB{
		client:     client,
		collection: coll,
	}, nil

}

type MongoDB struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (db *MongoDB) Stop(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}

func (db *MongoDB) Write(ctx context.Context, rawRecord any) error {
	_, err := db.collection.InsertOne(ctx, rawRecord)
	return err
}

func (db *MongoDB) read(ctx context.Context, filter bson.M) (ImageRecord, error) {
	var record MongoImageRecord
	err := db.collection.FindOne(ctx, filter).Decode(&record)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &record, err
}

func (db *MongoDB) FindByTags(ctx context.Context, tags []string) (ImageRecord, error) {
	filter := bson.M{"tags": bson.M{"$in": tags}}
	return db.read(ctx, filter)
}
