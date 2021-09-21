package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"dumpster/pkg/config"

	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DumpsterRepo - MongoDB Collection
type DumpsterRepo struct {
	*mongo.Collection
}

func connectLoop(ctx context.Context, client *options.ClientOptions) (*mongo.Client, error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeout := 5 * time.Minute

	timeoutExceeded := time.After(timeout)
	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %s timeout", timeout)

		case <-ticker.C:
			db, err := mongo.Connect(ctx, client)
			if err == nil {
				return db, nil
			}
			log.Println(errors.Wrapf(err, "Ticker Failed to connect to db %s", client.Hosts))
		}
	}
}

// NewDumpsterRepo - Connect to Database and return connection
func NewDumpsterRepo(ctx context.Context, config config.Config) (*DumpsterRepo, error) {
	clientOptions := options.Client().ApplyURI(config.URI)

	// TODO: Implement Initial Retry Logic Here Maybe? or higherlevel in main function?
	// Connect to MongoDB
	//client, err := mongo.Connect(ctx, clientOptions)
	client, err := connectLoop(ctx, clientOptions)

	if err != nil {
		return &DumpsterRepo{}, err
	}
	// Check the connection - Reduces Client resiliance...
	err = client.Ping(ctx, nil)

	if err != nil {
		return &DumpsterRepo{}, err
	}

	// Connect to the Configured Collection
	collection := client.Database(config.DB).Collection(config.Collection)

	return &DumpsterRepo{collection}, nil
}
