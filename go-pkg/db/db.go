package db

import (
	"context"
	"fmt"
	"github.com/benka-me/users/go-pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init(cfg *config.Config) (*mongo.Collection, *mongo.Client) {
	ctx := context.Background()
	uri := fmt.Sprintf("mongodb://%s:%s", cfg.DbHost, cfg.DbPort)
	credential := options.Credential{
		Username: cfg.DbUser,
		Password: cfg.DbPWD,
	}
	clientOpts := options.Client().ApplyURI(uri).SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		panic(err)
	}
	return client.Database(config.PgDatabase).Collection(config.PgCollection), client
}
