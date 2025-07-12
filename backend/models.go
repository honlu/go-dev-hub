package main

import (
	"github.com/lib/pq"
)

type Blog struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	ContentMD   string         `gorm:"type:text;not null" json:"content_md"`
	ContentHTML string         `gorm:"type:text" json:"content_html"`
	Category    string         `gorm:"type:varchar(50)" json:"category"`
	Tags        pq.StringArray `gorm:"type:text[]" json:"tags"`
	Status      string         `gorm:"type:varchar(20);default:'draft'" json:"status"`
	CreatedAt   int64          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   int64          `gorm:"autoUpdateTime" json:"updated_at"`
}

type Question struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Title      string         `gorm:"type:text;not null" json:"title"`
	Answer     string         `gorm:"type:text;not null" json:"answer"`
	Category   string         `gorm:"type:varchar(50)" json:"category"`
	Tags       pq.StringArray `gorm:"type:text[]" json:"tags"`
	Difficulty string         `gorm:"type:varchar(20)" json:"difficulty"`
	CreatedAt  int64          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  int64          `gorm:"autoUpdateTime" json:"updated_at"`
}
