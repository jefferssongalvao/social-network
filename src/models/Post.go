package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"author_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Likes      string    `json:"likes"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("title required")
	}
	if post.Content == "" {
		return errors.New("content required")
	}
	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}

func (post *Post) Prepare() error {
	if error := post.validate(); error != nil {
		return error
	}

	post.format()

	return nil
}
