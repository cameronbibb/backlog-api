package models

import "time"

type Game struct {
	ID           uint      `gorm:"primaryKey"`
	Title        string    `gorm:"not null" json:"title"`
	Description  string    `json:"description"`
	Genre        string    `gorm:"not null" json:"genre"`
	Rating       string    `json:"rating"`
	IGDBID       uint      `json:"igdb_id"`
	ThumbnailURL string    `json:"thumbnail_url"`
	CoverArtURL  string    `json:"cover_art_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
