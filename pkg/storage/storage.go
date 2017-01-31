package storage

import (
	"errors"
	"github.com/byoc-io/gutenberg/pkg/model"
)

var (
	// Limit default
	Limit = 100

	// ErrNotFound is the error returned by storages if a resource cannot be found.
	ErrNotFound = errors.New("not found")

	// ErrAlreadyExists is the error returned by storages if a resource is taken during a create.
	ErrAlreadyExists = errors.New("already exists")
)

// Repository is the database interface used by the server.
type Repository interface {
	ListPosts(pagination *Pagination) ([]*model.Post, error)
	GetPost(id string) (*model.Post, error)
	InsertPost(post *model.Post) error
	UpdatePost(post *model.Post) error
	HealthCheck() error
}

type Pagination struct {
	Limit  int
	Offset int
}

func NewPagination(perPage, page int) *Pagination {
	return &Pagination{
		Limit:  perPage,
		Offset: page * perPage,
	}
}
