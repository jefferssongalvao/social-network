package models

import "time"

type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   string    `json:"author_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Likes      string    `json:"likes"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}
