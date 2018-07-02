package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Session is to store informations about a recorded session
type Session struct {
	gorm.Model
	Name      string    `json:"session_name"`
	StartDate time.Time `json:"start_date"`
	Labels    []Label   `json:"labels"`
}

// BeforeDelete for labelset is used to clean up LabelTemplates
func (session *Session) BeforeDelete(tx *gorm.DB) (err error) {
	// Remove invoice notes:
	var labels []Label

	tx.Model(&session).Related(&labels)

	for _, item := range labels {
		tx.Delete(&item)
	}
	return
}

// Label is an annotated event
type Label struct {
	gorm.Model
	SessionID   uint      `json:"session_id"`
	Description string    `json:"description"`
	Subject     string    `json:"subject"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	CreatedBy   string    `json:"created_by"`
}

// LabelTemplate is one node in the hierarchy
type LabelTemplate struct {
	ID              uint            `gorm:"primary_key"`
	CreatedAt       time.Time       `json:"created_at,omitempty"`
	UpdatedAt       time.Time       `json:"updated_at,omitempty"`
	DeletedAt       *time.Time      `json:"deleted_at,omitempty"`
	Description     string          `json:"description"`
	LabelSetID      uint            `json:"labelset_id"`
	LabelTemplateID uint            `json:"parent_id"`
	Children        []LabelTemplate `json:"children"`
}

// LabelSet is a collaction of hierarchical labels
type LabelSet struct {
	gorm.Model
	Name   string          `json:"name"`
	Labels []LabelTemplate `json:"labels"`
}

// BeforeDelete for labelset is used to clean up LabelTemplates
func (labelset *LabelSet) BeforeDelete(tx *gorm.DB) (err error) {
	// Remove invoice notes:
	var labelTemps []LabelTemplate

	tx.Model(&labelset).Related(&labelTemps)

	for _, item := range labelTemps {
		tx.Delete(&item)
	}
	return
}
