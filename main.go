package main

import (
	"log"
	"net/http"
	"poll-app/controllers"
	"poll-app/initializers"
	"poll-app/routes"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	PollController       controllers.PollController
	PollAnswerController controllers.PollAnswerController
	PollRouteController  routes.PollRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

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

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	PollRouteController.PollRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}
