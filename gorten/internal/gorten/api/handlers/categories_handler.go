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

type CategoryHandlerImpl interface {
	List(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	UpdateByID(c *gin.Context)
}

type CategoryHandler struct {
	categoryService services.CategoryServiceImpl
}

func Category(s services.CategoryServiceImpl) *CategoryHandler {
	return &CategoryHandler{categoryService: s}
}

func (h *CategoryHandler) List(c *gin.Context) {
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

	categories, err := h.categoryService.List(c, paginationConfig.Offset(), paginationConfig.Limit(), sort)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
	categoryID := c.Param("id")
	category, err := h.categoryService.GetByID(c, categoryID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": category})
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var newCategory models.Category
	if err := c.BindJSON(&newCategory); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			logs.Logger.Printf("Error on validate fields on CategoryHandler::Create. Reason: %v", err)
			validation := utils.ValidationErrors(ve)
			_ = c.Error(validation)
			return
		}

		//The error from c.Error() matches the argument, so no need to check
		_ = c.Error(pkgerr.ErrInvalidRequestPayload)
		return
	}

	err := h.categoryService.Create(c, &newCategory)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"message":  http.StatusText(http.StatusCreated),
			"category": newCategory,
		},
	)
}

func (h *CategoryHandler) UpdateByID(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			logs.Logger.Printf("Error on validate fields on CategoryHandler::Update. Reason: %v", err)
			validation := utils.ValidationErrors(ve)
			_ = c.Error(validation)
			return
		}
		_ = c.Error(err)
		return
	}
	err := h.categoryService.UpdateByID(c, id, &category)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
