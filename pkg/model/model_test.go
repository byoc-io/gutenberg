package model

import (
	"github.com/byoc-io/gutenberg/pkg/util"
	"testing"
)

func TestNewPost(t *testing.T) {
	title := "The Go Programming Language"
	html := "<h1>content</h1>"
	authorID := util.NewUUID()

	p := NewPost(title, html, authorID)
	if p.ID == "" || len(p.ID) < 26 {
		t.Errorf("ID should not be null")
	}
}
