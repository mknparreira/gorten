package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gorten/internal/gorten/api/handlers"
	"gorten/internal/gorten/models"
	"gorten/test/factories"
	"gorten/test/integration/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupCompanyRouter(handler handlers.CompanyHandlerImpl) *gin.Engine {
	router := gin.Default()
	router.GET("/api/v1/companies", handler.List)
	router.GET("/api/v1/companies/:id", handler.GetByID)
	router.POST("/api/v1/companies", handler.Create)
	router.PUT("/api/v1/companies/:id", handler.UpdateByID)
	return router
}

func TestCompanyHandler_List(t *testing.T) {
	company := factories.CompanyFactory()
	expectedCompanies := []models.Company{*company}
	mockCompanyService := new(mocks.MockCompanyService)
	mockCompanyService.On("List", mock.Anything, 0, 10, "desc").Return(expectedCompanies, nil)

	companyHandler := handlers.Company(mockCompanyService)
	router := setupCompanyRouter(companyHandler)

	ctx := context.Background()
	req, _ := http.NewRequestWithContext(ctx, "GET", "/api/v1/companies", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), company.Name)
	mockCompanyService.AssertCalled(t, "List", mock.Anything, 0, 10, "desc")
}

func TestCompanyHandler_CompanyByID(t *testing.T) {
	company := factories.CompanyFactory()
	mockCompanyService := new(mocks.MockCompanyService)
	mockCompanyService.On("GetByID", mock.Anything, company.CompanyID).Return(company, nil)

	companyHandler := handlers.Company(mockCompanyService)
	router := setupCompanyRouter(companyHandler)

	ctx := context.Background()
	req, _ := http.NewRequestWithContext(ctx, "GET", "/api/v1/companies/"+company.CompanyID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), company.Name)
	mockCompanyService.AssertCalled(t, "GetByID", mock.Anything, company.CompanyID)
}

func TestCompanyHandler_Create(t *testing.T) {
	newCompany := factories.CompanyFactory(func(u *models.Company) {
		u.CompanyID = ""
		u.CreatedAt = time.Time{}
	})
	ctx := context.Background()
	mockCompanyService := new(mocks.MockCompanyService)
	mockCompanyService.On("Create", mock.Anything, newCompany).Return(nil)

	companyHandler := handlers.Company(mockCompanyService)
	router := setupCompanyRouter(companyHandler)

	body := map[string]string{
		"name":    newCompany.Name,
		"address": newCompany.Address,
		"contact": newCompany.Contact,
		"email":   newCompany.Email,
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req, _ := http.NewRequestWithContext(ctx, "POST", "/api/v1/companies", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockCompanyService.AssertCalled(t, "Create", mock.Anything, newCompany)
}

func TestCompanyHandler_UpdateByID(t *testing.T) {
	company := factories.CompanyFactory()
	ctx := context.Background()
	mockCompanyService := new(mocks.MockCompanyService)
	mockCompanyService.On("UpdateByID", mock.Anything, company.CompanyID, mock.AnythingOfType("*models.Company")).Return(nil)

	companyHandler := handlers.Company(mockCompanyService)
	router := setupCompanyRouter(companyHandler)

	body := map[string]string{
		"companyID": company.CompanyID,
		"name":      company.Name,
		"address":   company.Address,
		"contact":   company.Contact,
		"email":     company.Email,
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req, _ := http.NewRequestWithContext(ctx, "PUT", "/api/v1/companies/"+company.CompanyID, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockCompanyService.AssertCalled(t, "UpdateByID", mock.Anything, company.CompanyID, mock.AnythingOfType("*models.Company"))
}
