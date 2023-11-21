package models

import "github.com/google/uuid"

type Poll struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Question string    `gorm:"not null" json:"question,omitempty"`
}
