package model

import (
	"time"

	"fmt"
	"github.com/byoc-io/gutenberg/pkg/util"
	"regexp"
)

const (
	Internal Visibility = "internal"
	Public   Visibility = "public"

	Draft     Status = "draft"
	Scheduled Status = "scheduled"
	Published Status = "published"
)

// Metadata represents extra information about a model.
type Metadata struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Post is a read model for posting views.
type Post struct {
	ID    string `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Slug  string `json:"slug" binding:"required"`

	AMP    string `json:"amp"`
	HTML   string `json:"html"`
	Mobile string `json:"mobile"`

	Image         string `json:"image"`
	Featured      bool   `json:"featured"`
	FeaturedImage string `json:"featured_image"`
	Page          bool   `json:"page"`
	Locale        string `json:"locale"`
	Status        Status `json:"status"`
	*Metadata
	Visibility Visibility `json:"visibility"`

	AuthorID    string     `json:"author_id"`
	PublishedAt *time.Time `json:"published_at"`
	PublishedBy string     `json:"published_by"`
	CreatedAt   *time.Time `json:"created_at"`
	CreatedBy   string     `json:"created_by"`
	UpdatedAt   *time.Time `json:"updated_at"`
	UpdatedBy   string     `json:"updated_by"`
}

// URL to access a post.
func (p *Post) URL() (string, error) {
	return "", nil
}

func NewSlug(uuid string, title string) string {
	slug := util.Slugify(title, true)

	pattern, _ := regexp.Compile("^\\w+")
	sufix := pattern.FindString(uuid)

	return fmt.Sprintf("%s-%s", slug, sufix)
}

// New creates a new post.
func NewPost(title string, html string, authorID string) *Post {
	now := time.Now()
	id := util.NewUUID()

	return &Post{
		ID:         id,
		Title:      title,
		Slug:       NewSlug(id, title),
		HTML:       html,
		Featured:   false,
		Page:       false,
		Status:     Draft,
		Locale:     "en_US.UTF-8",
		Visibility: Public,
		AuthorID:   authorID,
		CreatedAt:  &now,
		UpdatedAt:  &now,
	}
}

// Tag model.
type Tag struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	Description string     `json:"description"`
	ParentID    string     `json:"parent_id"`
	Visibility  Visibility `json:"visibility"`
	*Metadata

	CreatedAt *time.Time `json:"created_at"`
	CreatedBy string     `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy string     `json:"updated_by"`
}

// User model.
type User struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Slug          string     `json:"slug"`
	Email         string     `json:"email"`
	ProfileImage  string     `json:"profile_image"`
	CoverImage    string     `json:"cover_image"`
	Bio           string     `json:"bio"`
	Website       string     `json:"website"`
	Location      string     `json:"location"`
	Facebook      string     `json:"facebook"`
	Twitter       string     `json:"twitter"`
	Accessibility string     `json:"accessibility"`
	Status        string     `json:"status"`
	Locale        string     `json:"locale"`
	Visibility    string     `json:"visibility"`
	Metadata      Metadata   `json:"Metadata"`
	Tour          string     `json:"tour"`
	Roles         []Role     `json:"roles"`
	LastLogin     *time.Time `json:"last_login"`

	CreatedAt *time.Time `json:"created_at"`
	CreatedBy *User      `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *User      `json:"updated_by"`
}

// Role model.
type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	CreatedAt *time.Time `json:"created_at"`
	CreatedBy User       `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy User       `json:"updated_by"`
}

type Permissions struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ActionType string `json:"action_type"`
	ObjectID   string `json:"object_id"`
	ObjectType string `json:"object_type"`

	CreatedAt *time.Time `json:"created_at"`
	CreatedBy User       `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy User       `json:"updated_by"`
}

// Status: draft, scheduled, published
type Status string

// Visibility: public, internal
type Visibility string
