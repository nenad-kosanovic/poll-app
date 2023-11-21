package models

import (
	"gorm.io/gorm"
)

type Poll struct {
	gorm.Model
	Question   string       `gorm:"not null" json:"question,omitempty" binding:"required"`
	PollAnswer []PollAnswer `binding:"min=2"`
}

type PollAnswer struct {
	gorm.Model
	Text   string `gorm:"not null" json:"text,omitempty" binding:"required"`
	Votes  int    `gorm:"not null" json:"votes"`
	PollId uint
}
