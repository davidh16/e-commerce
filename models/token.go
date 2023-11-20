package models

import "time"

type Token struct {
	Uuid      string    `json:"uuid" gorm:"unique;type:uuid; column:uuid;default:uuid_generate_v4()"`
	Token     string    `json:"token"`
	IsUsed    bool      `json:"is_used"`
	ExpiresAt time.Time `json:"expires_at"`
	UserUuid  string    `json:"user_uuid"`
}

func (m Token) TableName() string {
	return "tokens"
}
