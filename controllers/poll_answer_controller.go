package controllers

import (
	"fmt"
	"poll-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PollAnswerController struct {
	DB *gorm.DB
}

func NewPollAnswerController(DB *gorm.DB) PollAnswerController {
	return PollAnswerController{DB}
}

func (pc *PollAnswerController) AddVote(ctx *gin.Context) {
	pollAnswerId := ctx.Param("pollAnswerId")

	var pollAnswer models.PollAnswer
	result := pc.DB.Model(&pollAnswer).First(&pollAnswer, "id = ?", pollAnswerId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": fmt.Sprintf("No poll answer with provided id %s exists!", pollAnswerId)})
		return
	}

	pollAnswer.Votes++
	pc.DB.Model(&pollAnswer).Updates(pollAnswer)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": pollAnswer})
}
