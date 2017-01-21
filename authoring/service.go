package authoring

import (
	"errors"

	"github.com/byoc-io/gutenberg/pkg/model"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

type Service interface {
	GetPost(key string) (model.Post, error)
}

type service struct {
	repository model.Repository
}

func (s *service) GetPost(key string) (model.Post, error) {
	if key == "" {
		return model.Post{}, ErrInvalidArgument
	}

	p, err := s.repository.GetPost(key)
	if err != nil {
		return model.Post{}, err
	}

	return assemble(p), nil
}

func NewService(repository model.Repository) Service {
	return &service{
		repository: repository,
	}
}

func assemble(post *model.Post) model.Post {
	return model.Post{
		ID: string(post.ID),
	}
}
