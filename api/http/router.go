package http

import (
	"Auth_service-and-Api_gateway/api/config"
	"Auth_service-and-Api_gateway/api/http/handlers"
	"Auth_service-and-Api_gateway/api/pkg/grpc_client"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Config ...
type Config struct {
	GrpcClient *grpc_client.GrpcClient
	Cfg        config.Config
}

func New(cnf Config) *gin.Engine {
	router := gin.New()

	router.Static("/images", "./static/images")

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	handlerV1 := handlers.New(&handlers.HandlerV1Config{

		GrpcClient: cnf.GrpcClient,
	})

	apiV1 := router.Group("/v1")

	apiV1.Use( /*handlerV1.AuthMiddleware()*/ )
	{
		apiV1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "Api gateway"})
		})

		apiV1.POST("/user", handlerV1.CreateUser)

	}
	return router
}
