package drivers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"user-service/pkg/usecases"
)

const (
	serviceName = "warehouse-service"
)

// SetupRouter configure the gine router and returns a pointer to a gin engine which can be run.
func SetupRouter(createUser usecases.CreateUserInterface, userGetter usecases.UserGetter, userUpdater usecases.UpdateUserInterface, userGetterFromDB usecases.UserGetterDB, usersGetter usecases.UsersGetter) *gin.Engine {

	r := NewDefaultRouter(serviceName)

	// Configure routes
	v1 := r.Group("//v1")
	{
		v1.POST("/create-user/", usecases.CreateUser(createUser, userGetter))
		v1.POST("/update-user/", usecases.UpdateUser(userUpdater))
		v1.GET("/get-user/", usecases.GetUser(userGetterFromDB))
		v1.GET("/get-users/", usecases.GetUsers(usersGetter))
	}

	return r
}

func NewDefaultRouter(serviceNameVal string, additionalMiddleware ...gin.HandlerFunc) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	r.Use(additionalMiddleware...)

	// Declares a config variable, assigning its valid methods and headers
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "OPTIONS", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Forwarded-Authorization", "Strict-Transport-Security"}

	r.Use(cors.New(config))

	return r
}
