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

type ProductRepositoryImpl interface {
	GetAll(ctx context.Context, skip int, limit int, sort string) ([]models.Product, error)
	GetByID(ctx context.Context, productID string) (*models.Product, error)
	Create(ctx context.Context, product *models.Product) error
	Update(ctx context.Context, product *models.Product) error
}

type ProductRepository struct {
	collection *mongo.Collection
}

func ProductRepositoryInit(client *mongo.Client, ctg *config.AppConfig, collectionProductName string) *ProductRepository {
	collection := client.Database(ctg.Mongo.DBName).Collection(collectionProductName)
	return &ProductRepository{collection: collection}
}

func (r *ProductRepository) GetAll(ctx context.Context, skip int, limit int, sort string) ([]models.Product, error) {
	var products []models.Product
	s := int64(skip)
	l := int64(limit)
	sorting := bson.D{{Key: "createdAt", Value: utils.ConvertStringSortforInteger(sort)}}

	cursor, err := r.collection.Find(ctx, bson.M{},
		options.Find().SetSkip(s).SetLimit(l).SetSort(sorting))

	if err != nil {
		return nil, fmt.Errorf("could not fetch products: %w", err)
	}

	//Defer the execution until the GetAll() finishes,
	//ensuring that the cursor is closed after the iteration, regardless of errors.
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			fmt.Printf("error closing cursor: %v\n", err)
		}
	}()

	for cursor.Next(ctx) {
		var product models.Product
		//cursor.Decode(&product) similar to c.BindJSON() to merge object
		if err := cursor.Decode(&product); err != nil {
			return nil, fmt.Errorf("could not decode product: %w", err)
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return products, nil
}

func (r *ProductRepository) GetByID(ctx context.Context, productID string) (*models.Product, error) {
	var product models.Product
	filter := bson.M{"productId": productID}

	err := r.collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("product not found: %w", err)
		}
		return nil, fmt.Errorf("could not fetch product: %w", err)
	}

	return &product, nil
}

func (r *ProductRepository) Create(ctx context.Context, product *models.Product) error {
	product.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		return fmt.Errorf("could not create product: %w", err)
	}

	return nil
}

func (r *ProductRepository) Update(ctx context.Context, product *models.Product) error {
	filter := bson.M{"productId": product.ProductID}

	_, err := r.collection.ReplaceOne(ctx, filter, product)
	if err != nil {
		return fmt.Errorf("could not update product: %w", err)
	}

	return nil
}
