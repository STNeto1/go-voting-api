package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Voting struct {
	ID        string         `sql:"type:uuid;primary_key" json:"id"`
	Title     string         `json:"title"`
	Start     time.Time      `json:"start"`
	End       time.Time      `json:"end"`
	UserID    string         `json:"-"`
	User      User           `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Options   []VotingOption `json:"options"`
}

func (voting *Voting) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.New()
	voting.ID = id.String()
	return
}
