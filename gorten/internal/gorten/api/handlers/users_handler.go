package handlers

import (
	"errors"
	"gorten/internal/gorten/models"
	"gorten/internal/gorten/services"
	pkgerr "gorten/pkg/errors"
	"gorten/pkg/pagination"
	"gorten/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandlerImpl interface {
	List(c *gin.Context)
	UserByID(c *gin.Context)
	Create(c *gin.Context)
	UpdateByID(c *gin.Context)
}

type UserHandler struct {
	userService services.UserServiceImpl
}

func User(s services.UserServiceImpl) *UserHandler {
	return &UserHandler{userService: s}
}

func (h *UserHandler) List(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	paginationConfig, err := pagination.PaginationInit(page, limit)

	if err != nil {
		_ = c.Error(err)
		return
	}

	if err := paginationConfig.Validate(); err != nil {
		_ = c.Error(err)
		return
	}

	users, err := h.userService.List(c, paginationConfig.Offset(), paginationConfig.Limit())
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *UserHandler) UserByID(c *gin.Context) {
	userID := c.Param("id")
	user, err := h.userService.GetByID(c, userID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) Create(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			log.Printf("Error on validate fields on UserHandler::Create. Reason: %v", err)
			validation := utils.ValidationErrors(ve)
			_ = c.Error(validation)
			return
		}

		log.Printf("Error on UserHandler::Create. Reason: %v", err)
		//The error from c.Error() matches the argument, so no need to check
		_ = c.Error(pkgerr.ErrInvalidRequestPayload)
		return
	}

	err := h.userService.Create(c, &newUser)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"message": http.StatusText(http.StatusCreated),
			"user":    newUser,
		},
	)
}

func (h *UserHandler) UpdateByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			log.Printf("Error on validate fields on UserHandler::Update. Reason: %v", err)
			validation := utils.ValidationErrors(ve)
			_ = c.Error(validation)
			return
		}
		_ = c.Error(err)
		return
	}
	err := h.userService.UpdateByID(c, id, &user)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
