package models

import "time"

type RefreshToken struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	TokenHash string    `json:"-"`
	UserAgent *string   `json:"user_agent,omitempty"`
	IP        *string   `json:"ip,omitempty"`
	Revoked   bool      `json:"revoked"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}
