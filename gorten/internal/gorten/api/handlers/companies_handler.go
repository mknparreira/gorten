package handlers

import (
	"errors"
	"gorten/internal/gorten/models"
	"gorten/internal/gorten/services"
	pkgerr "gorten/pkg/errors"
	"gorten/pkg/logs"
	"gorten/pkg/pagination"
	"gorten/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CompanyHandlerImpl interface {
	List(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	UpdateByID(c *gin.Context)
}

type CompanyHandler struct {
	companyService services.CompanyServiceImpl
}

func Company(s services.CompanyServiceImpl) *CompanyHandler {
	return &CompanyHandler{companyService: s}
}

func (h *CompanyHandler) List(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	sort := c.DefaultQuery("sort", "desc")

	paginationConfig, err := pagination.PaginationInit(page, limit)

	if err != nil {
		_ = c.Error(err)
		return
	}

	if err := paginationConfig.Validate(); err != nil {
		_ = c.Error(err)
		return
	}

	companies, err := h.companyService.List(c, paginationConfig.Offset(), paginationConfig.Limit(), sort)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"companies": companies})
}

func (h *CompanyHandler) GetByID(c *gin.Context) {
	companyID := c.Param("id")
	company, err := h.companyService.GetByID(c, companyID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"company": company})
}

func (h *CompanyHandler) Create(c *gin.Context) {
	var newCompany models.Company
	if err := c.BindJSON(&newCompany); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			logs.Logger.Printf("Error on validate fields on CompanyHandler::Create. Reason: %v", err)
			validation := utils.ValidationErrors(ve)
			_ = c.Error(validation)
			return
		}

		//The error from c.Error() matches the argument, so no need to check
		_ = c.Error(pkgerr.ErrInvalidRequestPayload)
		return
	}

	err := h.companyService.Create(c, &newCompany)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"message": http.StatusText(http.StatusCreated),
			"company": newCompany,
		},
	)
}

func (h *CompanyHandler) UpdateByID(c *gin.Context) {
	id := c.Param("id")
	var company models.Company
	if err := c.BindJSON(&company); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			logs.Logger.Printf("Error on validate fields on CompanyHandler::Update. Reason: %v", err)
			validation := utils.ValidationErrors(ve)
			_ = c.Error(validation)
			return
		}
		_ = c.Error(err)
		return
	}
	err := h.companyService.UpdateByID(c, id, &company)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
