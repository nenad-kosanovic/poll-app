package controllers

import (
	"gin-api-mysql-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PollController struct {
	DB *gorm.DB
}

func NewPollController(DB *gorm.DB) PollController {
	return PollController{DB}
}

func (qc *PollController) CreatePoll(ctx *gin.Context) {
	var payload *models.Poll

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newPoll := models.Poll{
		Question:   payload.Question,
		PollAnswer: payload.PollAnswer,
	}

	result := qc.DB.Create(&newPoll)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newPoll})

}

func (qc *PollController) FindPolls(ctx *gin.Context) {

	var polls []models.Poll
	results := qc.DB.Model(&models.Poll{}).Preload("PollAnswer").Find(&polls)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(polls), "data": polls})
}

func (qc *PollController) FindPollById(ctx *gin.Context) {

	pollId := ctx.Param("pollId")

	var poll models.Poll
	result := qc.DB.Model(&poll).Preload("PollAnswer").First(&poll, "id = ?", pollId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No poll with that id exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": poll})
}
