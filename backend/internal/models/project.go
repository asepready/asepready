package models

import "time"

// Project represents a portfolio project
type Project struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Technologies []string  `json:"technologies"`
	GithubURL    string    `json:"github_url"`
	DemoURL      string    `json:"demo_url"`
	ImageURL     string    `json:"image_url"`
	IsFeatured   bool      `json:"is_featured"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// ProjectCreate represents project creation request
type ProjectCreate struct {
	Title        string   `json:"title" validate:"required,min=3,max=200"`
	Description  string   `json:"description"`
	Technologies []string `json:"technologies"`
	GithubURL    string   `json:"github_url" validate:"omitempty,url"`
	DemoURL      string   `json:"demo_url" validate:"omitempty,url"`
	ImageURL     string   `json:"image_url" validate:"omitempty,url"`
	IsFeatured   bool     `json:"is_featured"`
}

// ProjectUpdate represents project update request
type ProjectUpdate struct {
	Title        *string   `json:"title,omitempty" validate:"omitempty,min=3,max=200"`
	Description  *string   `json:"description,omitempty"`
	Technologies *[]string `json:"technologies,omitempty"`
	GithubURL    *string   `json:"github_url,omitempty" validate:"omitempty,url"`
	DemoURL      *string   `json:"demo_url,omitempty" validate:"omitempty,url"`
	ImageURL     *string   `json:"image_url,omitempty" validate:"omitempty,url"`
	IsFeatured   *bool     `json:"is_featured,omitempty"`
}
