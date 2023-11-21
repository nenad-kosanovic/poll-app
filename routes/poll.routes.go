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

func (qc *PollRouteController) PollRoute(rg *gin.RouterGroup) {

	router := rg.Group("polls")
	router.GET("", qc.pollController.FindPolls)
}
