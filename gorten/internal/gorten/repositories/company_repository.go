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

type CompanyRepositoryImpl interface {
	GetAll(ctx context.Context, skip int, limit int, sort string) ([]models.Company, error)
	GetByID(ctx context.Context, companyID string) (*models.Company, error)
	Create(ctx context.Context, company *models.Company) error
	Update(ctx context.Context, company *models.Company) error
}

type CompanyRepository struct {
	collection *mongo.Collection
}

func CompanyRepositoryInit(client *mongo.Client, ctg *config.AppConfig, collectionCompanyName string) *CompanyRepository {
	collection := client.Database(ctg.Mongo.DBName).Collection(collectionCompanyName)
	return &CompanyRepository{collection: collection}
}

func (r *CompanyRepository) GetAll(ctx context.Context, skip int, limit int, sort string) ([]models.Company, error) {
	var companies []models.Company
	s := int64(skip)
	l := int64(limit)
	sorting := bson.D{{Key: "createdAt", Value: utils.ConvertStringSortforInteger(sort)}}

	cursor, err := r.collection.Find(ctx, bson.M{},
		options.Find().SetSkip(s).SetLimit(l).SetSort(sorting))

	if err != nil {
		return nil, fmt.Errorf("could not fetch companies: %w", err)
	}

	//Defer the execution until the GetAll() finishes,
	//ensuring that the cursor is closed after the iteration, regardless of errors.
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			fmt.Printf("error closing cursor: %v\n", err)
		}
	}()

	for cursor.Next(ctx) {
		var company models.Company
		//cursor.Decode(&company) similar to c.BindJSON() to merge object
		if err := cursor.Decode(&company); err != nil {
			return nil, fmt.Errorf("could not decode company: %w", err)
		}
		companies = append(companies, company)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return companies, nil
}

func (r *CompanyRepository) GetByID(ctx context.Context, companyID string) (*models.Company, error) {
	var company models.Company
	filter := bson.M{"companyId": companyID}

	err := r.collection.FindOne(ctx, filter).Decode(&company)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("company not found: %w", err)
		}
		return nil, fmt.Errorf("could not fetch company: %w", err)
	}

	return &company, nil
}

func (r *CompanyRepository) Create(ctx context.Context, company *models.Company) error {
	company.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, company)
	if err != nil {
		return fmt.Errorf("could not create company: %w", err)
	}

	return nil
}

func (r *CompanyRepository) Update(ctx context.Context, company *models.Company) error {
	filter := bson.M{"companyId": company.CompanyID}

	_, err := r.collection.ReplaceOne(ctx, filter, company)
	if err != nil {
		return fmt.Errorf("could not update company: %w", err)
	}

	return nil
}
