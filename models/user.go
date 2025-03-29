package models

import "time"

type User struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	Username     string        `gorm:"unique;not null" json:"username"`
	Email        string        `gorm:"unique;not null" json:"email"`
	PasswordHash string        `gorm:"not null" json:"-"`
	Playlists    []Playlist    `json:"playlists"`
	BacklogItems []BacklogItem `json:"backlog_items"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}
