package constants

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")

	ErrWrongPassword  = errors.New("wrong password")
	ErrExpiredSession = errors.New("expired session")

	ErrInvalidToken       = errors.New("invalid token")
	ErrInvalidFingerprint = errors.New("invalid fingerprint")

	ErrNil = errors.New("no matching record found in redis database")
)