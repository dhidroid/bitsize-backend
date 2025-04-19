package utils

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var countersCollection *mongo.Collection

// InitializeCounter should be called once to initialize the counters collection
func InitializeCounter(collection *mongo.Collection) {
	countersCollection = collection
}

// GenerateUID generates a sequential numeric UID
func GenerateUID() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": "userid"}
	update := bson.M{"$inc": bson.M{"seq": 1}}

	var result struct {
		Seq int `bson:"seq"`
	}

	// Find and update the sequence counter
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	err := countersCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		return 0, fmt.Errorf("could not generate UID: %v", err)
	}

	// Return the incremented sequence number
	return result.Seq, nil
}
