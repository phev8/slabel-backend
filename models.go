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
	gorm.Model
	Description     string          `json:"description"`
	LabelTemplateID uint            `json:"parent_id"`
	Children        []LabelTemplate `json:"children"`
}

// LabelSet is a collaction of hierarchical labels
type LabelSet struct {
	gorm.Model
	Name   string          `json:"name"`
	Labels []LabelTemplate `json:"labels"`
}
