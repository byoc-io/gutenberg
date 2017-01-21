package model

import (
	"errors"
)

// ErrUnknown is used when a post could not be found.
var ErrUnknown = errors.New("unknown")

type Repository interface {
	GetPost(key string) (*Post, error)
	CreatePost(post *Post) (*Post, error)
}
