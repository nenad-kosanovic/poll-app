package models

import (
	"gorm.io/gorm"
)

type Poll struct {
	gorm.Model
	Question   string `gorm:"not null" json:"question,omitempty"`
	PollAnswer []PollAnswer
}

type PollAnswer struct {
	gorm.Model
	Text   string `gorm:"not null" json:"text,omitempty"`
	Votes  int    `gorm:"not null" json:"votes"`
	PollId uint
}
