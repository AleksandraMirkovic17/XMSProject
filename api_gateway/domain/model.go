package domain

import (
	"time"
)

type Comment struct {
	Id      string
	Content string
	Date    time.Time
	User    string
}

type Reaction struct {
	User     string
	Reaction int
}

type Post struct {
	Id        string
	User      string
	PostText  string
	Images    string
	Links     []string
	Date      time.Time
	Reactions []Reaction

	Comments  []Comment
	IsDeleted bool
}
