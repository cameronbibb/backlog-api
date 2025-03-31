package models

import "time"

type BacklogItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Status    string    `gorm:"not null" json:"status"`
	GameID    uint      `gorm:"not null" json:"game_id"`
	Game      Game      `gorm:"foreignKey:GameID" json:"game"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
