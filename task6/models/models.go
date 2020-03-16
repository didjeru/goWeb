package models

// Post Blog post
type Post struct {
	ID      string `bson:"_id,omitempty"`
	Title   string `bson:"title"`
	Content string `bson:"content"`
}
