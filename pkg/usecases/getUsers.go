package usecases

import (
	"github.com/gin-gonic/gin"
	"user-service/pkg/entities"
)

type getUsersResponse struct {
	Users []entities.User `json:"users"`
}

type UsersGetter interface {
	GetUsers() ([]entities.User, error)
}

func GetUsers(usersGetter UsersGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := usersGetter.GetUsers()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		response := getUsersResponse{
			Users: users,
		}
		c.JSON(200, response)
	}
}
