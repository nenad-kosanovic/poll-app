package routes

import (
	"poll-app/controllers"

	"github.com/gin-gonic/gin"
)

type PollRouteController struct {
	pollController controllers.PollController
	pollAnserController controllers.PollAnswerController
}

func NewPollRouteController(pollController controllers.PollController, pollAnserController controllers.PollAnswerController) PollRouteController {
	return PollRouteController{pollController,pollAnserController}
}

func (pc *PollRouteController) PollRoute(rg *gin.RouterGroup) {

	router := rg.Group("polls")
	router.GET("/:pollId", pc.pollController.FindPollById)
	router.GET("", pc.pollController.FindPolls)
	router.POST("", pc.pollController.CreatePoll)
	router.DELETE("/:pollId", pc.pollController.DeletePoll)
	router.PUT("/:pollId/answers/:pollAnswerId", pc.pollAnserController.AddVote)
}
