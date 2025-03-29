package models

import "time"

type Playlist struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	Name         string        `gorm:"not null" json:"name"`
	Description  string        `json:"description"`
	UserId       uint          `json:"user_id"`
	BacklogItems []BacklogItem `gorm:"many2many:playlist_backlog_items" json:"backlog_items"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}
