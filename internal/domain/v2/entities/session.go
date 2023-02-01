package entity 

import "time"

type Session struct {
	ID          string
	UserID      string
	Fingerprint string
	AccessToken string
	UpdateToken string

	// ..
	CreatedAt time.Time
	ExpiresAt time.Time
}
