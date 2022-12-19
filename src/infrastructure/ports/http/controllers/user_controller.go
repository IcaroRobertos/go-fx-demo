package controllers

import (
	"net/http"

	"github.com/IcaroRobertos/go-fx-demo/src/application/usecases"
	"github.com/IcaroRobertos/go-fx-demo/src/domain/entities"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCases *usecases.UserUseCases
}

func NewUserController(userUseCases *usecases.UserUseCases) *UserController {
	return &UserController{
		UserUseCases: userUseCases,
	}
}

func (uc *UserController) Create(c *gin.Context) {
	var body entities.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.UserUseCases.Create(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
