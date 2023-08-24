package mongodb

import (
	"MedodsProject/pkg/config"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, cfg config.MongoDB) (db *mongo.Database, err error) {
	// var mongoDBURL string
	// var isAuth bool
	// if username == "" && password == "" {
	// 	mongoDBURL = fmt.Sprintf("mongodb://%s:%s", host, port)
	// } else {
	// 	isAuth = true
	// 	mongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	// }

	mongoDBURL := fmt.Sprintf("mongodb://%s:%d", cfg.Host, cfg.Port)

	clientOptions := options.Client().ApplyURI(mongoDBURL)
	// if isAuth {
	// 	clientOptions.SetAuth(options.Credential{
	// 		AuthSource:	authDB,
	// 		Username:	username,
	// 		Password:	password,
	// 	})
	// }

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongoDB due to error: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongoDB due to error: %v", err)
	}

	return client.Database(cfg.DBName), nil
}
