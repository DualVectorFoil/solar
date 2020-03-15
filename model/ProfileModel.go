package model

import "github.com/jinzhu/gorm"

type Profile struct {
	gorm.Model
	AccountID           string `gorm:"not null;index:idx_account_id"`
	PhoneNum     string
	UserName     string `gorm:"not null"`
	Email        string
	AvatarUrl    string
	Pwd          string `gorm:"not null"`
	Locale       string
	Bio          string
	Followers    int
	Following    int
	ArtworkCount int
	RegisterAt   int64 `gorm:"not null"`
	LastLoginAt  int64 `gorm:"not null"`
}
