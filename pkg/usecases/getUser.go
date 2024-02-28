package usecases

import (
	"github.com/gin-gonic/gin"
	"user-service/pkg/entities"
)

type UserGetterDB interface {
	GetUserFromDB(id string) (entities.User, error)
}
type getUserRequest struct {
	Id string `json:"id" binding:"required"`
}

type getUserResponse struct {
	entities.User `json:"user"`
}

func GetUser(userGetterDB UserGetterDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request getUserRequest
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(400, gin.H{"error": "Bad request"})
			return
		}
		user, err := userGetterDB.GetUserFromDB(request.Id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		response := getUserResponse{
			User: user,
		}
		c.JSON(200, response)
	}
}
