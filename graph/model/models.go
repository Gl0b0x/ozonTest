package model

import "time"

type Comment struct {
	ID        string    `json:"id"`
	PostID    string    `json:"postId"`
	ParentID  *string   `json:"parentId,omitempty"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	Author    *User     `json:"author"`
}

type Mutation struct {
}

type Post struct {
	ID              string     `json:"id"`
	Title           string     `json:"title"`
	Content         string     `json:"content"`
	Comments        []*Comment `json:"comments,omitempty"`
	CommentsEnabled bool       `json:"commentsEnabled"`
	CreatedAt       time.Time  `json:"createdAt"`
	Author          *User      `json:"author"`
}

type Query struct {
}

type Subscription struct {
}

type User struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Posts    []*Post `json:"posts,omitempty"`
}
