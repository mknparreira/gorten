package repositories

import (
	"context"
	"errors"
	"fmt"
	"gorten/internal/gorten/config"
	"gorten/internal/gorten/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl interface {
	GetAll(ctx context.Context) ([]models.User, error)
	GetByID(ctx context.Context, userID string) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
}

type UserRepository struct {
	collection *mongo.Collection
}

func UserRepositoryInit(client *mongo.Client, ctg *config.AppConfig, collectionName string) *UserRepository {
	collection := client.Database(ctg.Mongo.DBName).Collection(collectionName)
	return &UserRepository{collection: collection}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("could not fetch users: %w", err)
	}

	//Defer the execution until the GetAll() finishes,
	//ensuring that the cursor is closed after the iteration, regardless of errors.
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			fmt.Printf("error closing cursor: %v\n", err)
		}
	}()

	for cursor.Next(ctx) {
		var user models.User
		//cursor.Decode(&user) similar to c.BindJSON() to merge object
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

func (r *UserRepository) GetByID(ctx context.Context, userID string) (*models.User, error) {
	var user models.User
	filter := bson.M{"userId": userID}

	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("could not fetch user: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}

	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	filter := bson.M{"userId": user.UserID}

	_, err := r.collection.ReplaceOne(ctx, filter, user)
	if err != nil {
		return fmt.Errorf("could not update user: %w", err)
	}

	return nil
}
