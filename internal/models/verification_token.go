package models

import "time"

type VerificationToken struct {
	ID     int32 `json:"id"`
	UserID int32 `json:"user_id"`
	// TokenHash is intentionally not exposed in JSON
	TokenHash string    `json:"-"`
	Kind      string    `json:"kind"` // "email_verify" | "password_reset"
	ExpiresAt time.Time `json:"expires_at"`
	Used      bool      `json:"used"`
	CreatedAt time.Time `json:"created_at"`
}
