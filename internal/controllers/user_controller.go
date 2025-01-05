package controllers

import (
	"net/http"
	"webserver/internal/models"
	"webserver/internal/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *services.UserService
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) ListUsers(ctx *gin.Context) {
	users, err := c.Service.ListUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
