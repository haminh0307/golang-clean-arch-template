package restapi

import (
	"net/http"

	_ "github.com/haminh0307/golang-clean-arch-template/internal/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Backend API
//	@version		1.0
//	@description	This is a sample server.
//	@termsOfService	http://swagger.io/terms/

//	@BasePath	/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@query.collection.format	multi

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.

func NewServer(host string, port string, controllers ...Controller) *http.Server {
	engine := gin.Default()

	// swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// health check
	engine.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })

	// configs cors
	engine.Use(CORSMiddleware())

	// api
	apiGroup := engine.Group("/")
	for _, controller := range controllers {
		controller.RegisterRoutes(apiGroup)
	}

	return &http.Server{
		Addr:    host + ":" + port,
		Handler: engine,
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)

			return
		}

		c.Next()
	}
}
