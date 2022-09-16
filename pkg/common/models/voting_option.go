package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VotingOption struct {
	ID          string `sql:"type:uuid;primary_key" json:"id"`
	Description string `json:"description"`
	Votes       int    `json:"votes"`
	VotingID    string `json:"-"`
	Voting      Voting `json:"-"`
}

func (voting *VotingOption) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.New()
	voting.ID = id.String()
	return
}
