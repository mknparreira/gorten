package repositories

import (
	"context"
	"errors"
	"fmt"
	"gorten/internal/gorten/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func UserRepositoryInit(client *mongo.Client, dbName, collectionName string) *UserRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &UserRepository{collection: collection}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("could not fetch users: %w", err)
	}

	defer func() {
		if err := cursor.Close(context.TODO()); err != nil {
			fmt.Printf("error closing cursor: %v\n", err)
		}
	}()

	for cursor.Next(context.TODO()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, fmt.Errorf("could not decode user: %w", err)
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return users, nil
}

func (r *UserRepository) GetByID(id string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("could not fetch user: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) Create(user *models.User) error {
	_, err := r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

func (r *UserRepository) Update(user *models.User) error {
	_, err := r.collection.ReplaceOne(context.TODO(), bson.M{"_id": user.UserID}, user)
	if err != nil {
		return fmt.Errorf("could not update user: %w", err)
	}
	return nil
}
