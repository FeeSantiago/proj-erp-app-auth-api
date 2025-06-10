package controllers

import (
	"context"
	"net/http"
	"proj-erp-auth/helpers"
	"time"

	"proj-erp-auth/repository"

	"github.com/gin-gonic/gin"
)

func GetUser(repo *repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")

		if err := helpers.MatchUserTypeToUserId(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		user, err := repo.GetByID(userId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
