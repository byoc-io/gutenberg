package model

// Tag model.
type Tag struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Metadata    Metadata `json:"metadata"`
	Name        string   `json:"id"`
	Parent      *Tag     `json:"tag"`
	Slug        string   `json:"id"`
}
