package model

import (
	"github.com/byoc-io/gutenberg/pkg/util"
	"time"
)

// Post is a read model for posting views.
type Post struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Slug     string `json:"slug"`
	Content  string `json:"content"`
	Image    string `json:"image"`
	Featured bool   `json:"featured"`
	Page     bool   `json:"page"`

	// Status: draft, scheduled, published
	Status   string `json:"status"`
	Language string `json:"language"`

	// Visibility: public, internal
	Visibility  string    `json:"visibility"`
	Metadata    Metadata  `json:"metadata"`
	Author      User      `json:"author"`
	PublishedAt time.Time `json:"published_at"`
	PublishedBy User      `json:"published_by"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   User      `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   User      `json:"updated_by"`
}

// URL to access a post.
func (p *Post) URL() (string, error) {
	return "", nil
}

// New creates a new post.
func NewPost(title string, content string, author *User) *Post {
	now := time.Now()
	return &Post{
		ID:         util.NewUUID(),
		Title:      title,
		Slug:       util.Slugify(title, true),
		Content:    content,
		Featured:   false,
		Page:       false,
		Status:     "draft",
		Language:   "en_US.UTF-8",
		Visibility: "public",
		Author:     *author,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}
