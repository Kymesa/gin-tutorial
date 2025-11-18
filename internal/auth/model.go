package auth

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
}

type RefreshToken struct {
	ID        uint   `gorm:"primaryKey"`
	Token     string `gorm:"uniqueIndex"`
	UserID    uint   `gorm:"index"`
	ExpiresAt time.Time
	Revoked   bool
	CreatedAt time.Time
}
