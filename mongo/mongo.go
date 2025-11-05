package mongo

import (
	"base_app/config"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo wraps a connected mongo client and database handle.
type Mongo struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var mongoInstance *Mongo

// InitMongo initializes a singleton Mongo connection using provided config.
// Safe to call multiple times; the first successful call wins.
func InitMongo() (*Mongo, error) {
	if mongoInstance != nil {
		return mongoInstance, nil
	}

	mongoConf := config.GetMongoConf()
	connectTimeout := mongoConf.ConnectTimeout

	if connectTimeout <= 0 {
		connectTimeout = 10
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(connectTimeout)*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(mongoConf.Uri)
	if mongoConf.Username != "" || mongoConf.Password != "" {
		clientOpts.SetAuth(options.Credential{
			Username:   mongoConf.Username,
			Password:   mongoConf.Password,
			AuthSource: mongoConf.AuthSource,
		})
	}
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}
	// Ping to verify connectivity
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	mongoInstance = &Mongo{
		Client:   client,
		Database: client.Database(mongoConf.Database),
	}
	return mongoInstance, nil
}

// GetMongo returns the initialized Mongo singleton, or nil if not initialized.
func GetMongo() *Mongo {
	return mongoInstance
}

// Close terminates the underlying Mongo client.
func (m *Mongo) Close(ctx context.Context) error {
	if m == nil || m.Client == nil {
		return nil
	}
	return m.Client.Disconnect(ctx)
}
