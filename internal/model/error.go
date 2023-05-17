package model

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrRoleInUse    = errors.New("role is in use")
)
