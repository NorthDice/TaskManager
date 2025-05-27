package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

type AuthMongo struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewAuthMongo() *AuthMongo {
	return &AuthMongo{
		// Initialize MongoDB client or session here if needed
	}
}
