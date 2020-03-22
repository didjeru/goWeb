package models

// Post Blog post
type Post struct {
	ID      int
	Title   string
	Content string
}

// NewPost Create a new post
func NewPost(id int, title string, content string) Post {
	return Post{
		ID:      id,
		Title:   title,
		Content: content,
	}
}
