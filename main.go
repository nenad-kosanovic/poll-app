package main

import (
	"poll-app/controllers"
	"poll-app/initializers"
	"poll-app/routes"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	PollController      controllers.PollController
	PollAnswerController controllers.PollAnswerController
	PollRouteController routes.PollRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	log.Println(config.HelloMessage)
	initializers.ConnectDB(&config)

	PollController = controllers.NewPollController(initializers.DB)
	PollAnswerController = controllers.NewPollAnswerController(initializers.DB)
	PollRouteController = routes.NewPollRouteController(PollController, PollAnswerController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	PollRouteController.PollRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}
