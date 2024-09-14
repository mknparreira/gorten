package repositories

import (
	"context"
	"errors"
	"fmt"
	"gorten/internal/gorten/config"
	"gorten/internal/gorten/models"
	"gorten/pkg/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CategoryRepositoryImpl interface {
	GetAll(ctx context.Context, skip int, limit int, sort string) ([]models.Category, error)
	GetByID(ctx context.Context, categoryID string) (*models.Category, error)
	Create(ctx context.Context, category *models.Category) error
	Update(ctx context.Context, category *models.Category) error
}

type CategoryRepository struct {
	collection *mongo.Collection
}

func CategoryRepositoryInit(client *mongo.Client, ctg *config.AppConfig, collectionCategoryName string) *CategoryRepository {
	collection := client.Database(ctg.Mongo.DBName).Collection(collectionCategoryName)
	return &CategoryRepository{collection: collection}
}

func (r *CategoryRepository) GetAll(ctx context.Context, skip int, limit int, sort string) ([]models.Category, error) {
	var categories []models.Category
	s := int64(skip)
	l := int64(limit)
	sorting := bson.D{{Key: "createdAt", Value: utils.ConvertStringSortforInteger(sort)}}

	cursor, err := r.collection.Find(ctx, bson.M{},
		options.Find().SetSkip(s).SetLimit(l).SetSort(sorting))

	if err != nil {
		return nil, fmt.Errorf("could not fetch categories: %w", err)
	}

	//Defer the execution until the GetAll() finishes,
	//ensuring that the cursor is closed after the iteration, regardless of errors.
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			fmt.Printf("error closing cursor: %v\n", err)
		}
	}()

	for cursor.Next(ctx) {
		var category models.Category
		//cursor.Decode(&category) similar to c.BindJSON() to merge object
		if err := cursor.Decode(&category); err != nil {
			return nil, fmt.Errorf("could not decode category: %w", err)
		}
		categories = append(categories, category)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return categories, nil
}

func (r *CategoryRepository) GetByID(ctx context.Context, categoryID string) (*models.Category, error) {
	var category models.Category
	filter := bson.M{"categoryId": categoryID}

	err := r.collection.FindOne(ctx, filter).Decode(&category)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("category not found: %w", err)
		}
		return nil, fmt.Errorf("could not fetch category: %w", err)
	}

	return &category, nil
}

func (r *CategoryRepository) Create(ctx context.Context, category *models.Category) error {
	category.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, category)
	if err != nil {
		return fmt.Errorf("could not create category: %w", err)
	}

	return nil
}

func (r *CategoryRepository) Update(ctx context.Context, category *models.Category) error {
	filter := bson.M{"categoryId": category.CategoryID}

	_, err := r.collection.ReplaceOne(ctx, filter, category)
	if err != nil {
		return fmt.Errorf("could not update category: %w", err)
	}

	return nil
}
