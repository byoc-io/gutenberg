package model

import (
	"github.com/byoc-io/gutenberg/pkg/util"
	"reflect"
	"testing"
	"time"
)

func TestConstruction(t *testing.T) {
	title := "The Go Programming Language"
	content := "<h1>content</h1>"
	author := &User{
		ID: util.NewUUID(),
	}

	p := NewPost(title, content, author)
	if p.ID == "" || len(p.ID) < 26 {
		t.Errorf("ID should not be null")
	}

	if p.Title != title {
		t.Errorf("Title = %v; want = %v", p.Title, title)
	}

	slug := "the-go-programming-language"
	if p.Slug != slug {
		t.Errorf("Slug = %v; want = %v", p.Slug, slug)
	}

	if p.Content != content {
		t.Errorf("Content = %v; want = %v", p.Content, content)
	}

	if p.Featured {
		t.Errorf("Featured = %v; want = %v", p.Featured, false)
	}

	if p.Page {
		t.Errorf("Page = %v; want = %v", p.Page, false)
	}

	if p.Status != "draft" {
		t.Errorf("Status = %v; want = %v", p.Status, "draft")
	}

	if p.Language != "en_US.UTF-8" {
		t.Errorf("Language = %v; want = %v", p.Language, "en_US.UTF-8")
	}

	if p.Author.ID != author.ID {
		t.Errorf("Author = %v; want = %v", p.Author, author)
	}

	if !p.PublishedAt.IsZero() {
		t.Errorf("PublishedAt = %v should be zero", p.CreatedAt)
	}

	user := User{}
	if !reflect.DeepEqual(p.PublishedBy, user) {
		t.Errorf("Published by = %v should be nil", p.PublishedBy)
	}

	now := time.Now()
	if p.CreatedAt.IsZero() || p.CreatedAt.After(now) {
		t.Errorf("CreatedAt = %v should be initialized before (%v)", p.CreatedAt, now)
	}

	if p.UpdatedAt.IsZero() || p.UpdatedAt.After(now) {
		t.Errorf("UpdatedAt = %v should be initialized before (%v)", p.CreatedAt, now)
	}
}
