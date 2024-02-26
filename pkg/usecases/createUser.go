package usecases

import (
	"context"
	"github.com/gin-gonic/gin"
	"user-service/pkg/entities"
)

type CreateUserRequest struct {
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Admin *bool  `json:"admin" binding:"required"`
}
type CreateUserResponse struct {
	Message string `json:"message"`
}

type CreateUserInterface interface {
	CreateUser(ctx context.Context, id string, email string, name string, admin bool) error
}
type UserGetter interface {
	GetUser(email string) (entities.User, error)
}

func CreateUser(CreateUserInterface CreateUserInterface, UserGetter UserGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request CreateUserRequest
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(400, "error bad request")
			return
		}

		email := request.Email
		name := request.Name
		admin := *request.Admin

		user, err := UserGetter.GetUser(email)

		err = CreateUserInterface.CreateUser(c, user.ID, email, name, admin)
		if err != nil {
			println(err.Error())
			if err.Error() == "pq: duplicate key value violates unique constraint \"users_user_email_key\"" {
				c.JSON(400, gin.H{"error": "user already exists"})
				return

			}
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "User created"})
	}
}
