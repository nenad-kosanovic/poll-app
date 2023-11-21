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

func (qc *PollController) FindPolls(ctx *gin.Context) {

	var polls []models.Poll
	results := qc.DB.Find(&polls)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(polls), "data": polls})
}
