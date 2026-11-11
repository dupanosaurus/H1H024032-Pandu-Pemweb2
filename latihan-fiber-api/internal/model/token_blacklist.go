package model

import "time"

type TokenBlacklist struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	TokenHash string    `gorm:"size:64;uniqueIndex" json:"-"`
	ExpiresAt time.Time `json:"kedaluwarsa_pada"`
	CreatedAt time.Time `json:"dibuat_pada"`
}

func (TokenBlacklist) TableName() string {
	return "token_blacklists"
}
