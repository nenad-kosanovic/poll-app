package routes

import (
	"gin-api-mysql-crud/controllers"

	"github.com/gin-gonic/gin"
)

type PollRouteController struct {
	pollController controllers.PollController
}

func NewPollRouteController(pollController controllers.PollController) PollRouteController {
	return PollRouteController{pollController}
}

func (pc *PollRouteController) PollRoute(rg *gin.RouterGroup) {

	router := rg.Group("polls")
	router.GET("/:pollId", pc.pollController.FindPollById)
	router.GET("", pc.pollController.FindPolls)
	router.POST("", pc.pollController.CreatePoll)
	router.DELETE("/:pollId", pc.pollController.DeletePoll)
}
