package auth

import "errors"

// ErrCheckpoint is returned when LinkedIn triggers a security checkpoint
var ErrCheckpoint = errors.New("linkedin checkpoint detected")
