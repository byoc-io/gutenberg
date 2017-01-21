package authoring

import (
	"github.com/byoc-io/gutenberg/pkg/model"
	"github.com/byoc-io/gutenberg/pkg/storage/memory"
	"testing"
)

func TestGetPost(t *testing.T) {
	r := memory.NewRepository()
	s := NewService(r)

	p1 := model.NewPost("title", "content", &model.User{})
	r.CreatePost(p1)
	_, err := s.GetPost(p1.ID)
	if err != nil {
		t.Fatal(err)
	}

	if _, error := s.GetPost(""); error != ErrInvalidArgument {
		t.Errorf("error = %v; expected = %v", error, ErrInvalidArgument)
	}
}
