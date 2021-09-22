package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"dumpster/pkg/storage"
)

// SaveObject - Save Reponse to database
func (d *DumpsterRepo) SaveObject(ctx context.Context, r *storage.Payload) error {
	// Set ID
	r.ID = primitive.NewObjectID()
	// Set time
	r.Date = time.Now()

	insertResult, err := d.InsertOne(ctx, r)
	if err != nil {
		return err
	}
	fmt.Printf("Inserted: %v", insertResult.InsertedID)
	return nil
}

// DeleteObject - Delete a Response
func (d *DumpsterRepo) DeleteObject(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	_, err = d.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil

}

// ListObjects - Return a List of Responses
func (d *DumpsterRepo) ListObjects(ctx context.Context) ([]*storage.Payload, error) {
	findOptions := options.Find()
	findOptions.SetLimit(50)

	var results []*storage.Payload

	cur, err := d.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return results, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {

		// create a value into which the single document can be decoded
		var elem storage.Payload
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Printf("Error Decoding Element: %v", err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		fmt.Printf("Error in Cursor: %v\n", err)
	}

	// Close the cursor once finished
	cur.Close(ctx)
	return results, nil
}

// GetByID - Return a reponse based ID given
func (d *DumpsterRepo) GetByID(ctx context.Context, id string) (*storage.Payload, error) {
	var result storage.Payload

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &storage.Payload{}, err
	}

	filter := bson.M{"_id": objID}
	err = d.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return &storage.Payload{}, err
	}

	return &result, nil
}

// GetByTwID - Return a reponse based ID given
func (d *DumpsterRepo) GetByTwID(ctx context.Context, id string) (*storage.Payload, error) {
	var result storage.Payload

	filter := bson.M{"twitter_id": id}
	err := d.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return &storage.Payload{}, err
	}

	return &result, nil
}
