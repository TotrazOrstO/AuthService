package user

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collection = "refresh_tokens"

type refreshTokenDB struct {
	collection *mongo.Collection
}

func (d *refreshTokenDB) SaveRefreshToken(ctx context.Context, userID, hashedRefreshToken string) error {
	result, err := d.collection.UpdateOne(context.Background(), bson.M{"_id": userID}, bson.M{"$set": bson.M{"refresh_token": hashedRefreshToken}})
	if err != nil {
		return fmt.Errorf("error updating refresh token: %w", err)
	}

	if result.ModifiedCount == 0 {
		log.Print("refresh token update did not modify any documents. trying to create new...")
		_, err := d.collection.InsertOne(ctx, RefreshToken{ID: userID, RefreshToken: hashedRefreshToken})
		if err != nil {
			return fmt.Errorf("insert one: %w", err)
		}
	}

	return nil
}

func (d *refreshTokenDB) RefreshTokenByID(ctx context.Context, userId string) (string, error) {
	var rt RefreshToken
	err := d.collection.FindOne(context.Background(), bson.M{"_id": userId}).Decode(&rt)
	if err != nil {
		return "", fmt.Errorf("findone: %w", err)
	}

	return rt.RefreshToken, nil
}

func NewRefreshTokenDB(database *mongo.Database) *refreshTokenDB {
	return &refreshTokenDB{collection: database.Collection(collection)}
}
