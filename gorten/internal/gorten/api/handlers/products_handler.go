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

type ProductHandlerImpl interface {
	List(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	UpdateByID(c *gin.Context)
}

type ProductHandler struct {
	productService services.ProductServiceImpl
}

func Product(s services.ProductServiceImpl) *ProductHandler {
	return &ProductHandler{productService: s}
}

func (h *ProductHandler) List(c *gin.Context) {
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

	products, err := h.productService.List(c, paginationConfig.Offset(), paginationConfig.Limit(), sort)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	productID := c.Param("id")
	product, err := h.productService.GetByID(c, productID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func (h *ProductHandler) Create(c *gin.Context) {
	var newProduct models.Product
	if err := c.BindJSON(&newProduct); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			logs.Logger.Printf("Error on validate fields on ProductHandler::Create. Reason: %v", err)
			validation := utils.ValidationErrors(ve)
			_ = c.Error(validation)
			return
		}

		//The error from c.Error() matches the argument, so no need to check
		_ = c.Error(pkgerr.ErrInvalidRequestPayload)
		return
	}

	err := h.productService.Create(c, &newProduct)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"message": http.StatusText(http.StatusCreated),
			"product": newProduct,
		},
	)
}

func (h *ProductHandler) UpdateByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			logs.Logger.Printf("Error on validate fields on ProductHandler::Update. Reason: %v", err)
			validation := utils.ValidationErrors(ve)
			_ = c.Error(validation)
			return
		}
		_ = c.Error(err)
		return
	}
	err := h.productService.UpdateByID(c, id, &product)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
