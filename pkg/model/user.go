package model

import (
	"time"
)

// User model.
type User struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	AccessToken   string    `json:"access_token"`
	Password      string    `json:"password"`
	Email         string    `json:"email"`
	Image         string    `json:"image"`
	Cover         string    `json:"cover"`
	Bio           string    `json:"bio"`
	Website       string    `json:"website"`
	Location      string    `json:"location"`
	Facebook      string    `json:"facebook"`
	Twitter       string    `json:"twitter"`
	Accessibility string    `json:"accessibility"`
	Status        string    `json:"status"`
	Language      string    `json:"language"`
	Visibility    string    `json:"visibility"`
	Metadata      Metadata  `json:"Metadata"`
	Tour          string    `json:"tour"`
	Roles         []Role    `json:"roles"`
	LastLogin     time.Time `json:"last_login"`
	CreatedAt     time.Time `json:"created_at"`
	CreatedBy     *User     `json:"created_by"`
	UpdatedAt     time.Time `json:"updated_at"`
	UpdatedBy     *User     `json:"updated_by"`
}

// Role model.
type Role struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	User        User      `json:"user"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   User      `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   User      `json:"updated_by"`
}
