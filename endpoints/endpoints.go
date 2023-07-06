package endpoints

import (
	"database/sql"
	"depogunabangunan/apps/handler"
	"depogunabangunan/apps/interfaces"

	"github.com/gin-gonic/gin"
)

// InitializeEndpoints initializes all the endpoints
func InitializeEndpoints(router *gin.RouterGroup, db *sql.DB,
	userRepository interfaces.IUserService,
	productRepository interfaces.IProductService) {

	userHandler := handler.NewUserHandler(userRepository)
	productHandler := handler.NewProductHandler(productRepository)

	// Define the user endpoints
	userEndpoints := router.Group("/users")
	{
		userEndpoints.GET("/getById/:id", userHandler.GetUserByID)
		userEndpoints.GET("/getByUsername/:username", userHandler.GetUserByUsername)
		userEndpoints.POST("/createUser", userHandler.CreateUser)
		userEndpoints.PUT("/updateUser/:id", userHandler.UpdateUser)
		userEndpoints.DELETE("/deleteUser/:id", userHandler.DeleteUser)
	}

	// Define the product endpoints
	productEndpoints := router.Group("/products")
	{
		productEndpoints.GET("/:id", productHandler.GetProductByID)
		productEndpoints.POST("/", productHandler.CreateProduct)
		productEndpoints.PUT("/:id", productHandler.UpdateProduct)
		productEndpoints.DELETE("/:id", productHandler.DeleteProduct)
	}

	// Define the login endpoint
	loginHandler := handler.NewLoginHandler(db, userRepository)
	router.POST("/login", loginHandler.Login)
}
