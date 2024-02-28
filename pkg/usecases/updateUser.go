package usecases

import (
	"context"
	"github.com/gin-gonic/gin"
)

type updateUserRequest struct {
	Id       string `json:"id" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Admin    bool   `json:"admin" binding:"required"`
	Disabled bool   `json:"disabled" binding:"required"`
}

type UpdateUserInterface interface {
	UpdateUser(ctx context.Context, id string, email string, name string, admin bool, disabled bool) error
}

func UpdateUser(UpdateUserInterface UpdateUserInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request updateUserRequest
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(400, "error bad request")
			return
		}

		err = UpdateUserInterface.UpdateUser(c, request.Id, request.Email, request.Name, request.Admin, request.Disabled)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "User updated"})
	}
}
