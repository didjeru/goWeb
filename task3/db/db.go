package db

import (
	"../models"
	"sync"
)

// DB database of project
type DB struct {
	posts map[int]models.Post
	mux   sync.Mutex
}

// AddNewPost add a new post
func (db *DB) AddNewPost(title string, content string) {
	id := db.GetCountPosts()
	db.posts[id] = models.NewPost(id, title, content)
}

// IsPostExists Check if post is exist
func (db *DB) IsPostExists(id int) bool {
	db.mux.Lock()
	_, ok := db.posts[id]
	db.mux.Unlock()
	return ok
}

// UpdatePost update an exists post
func (db *DB) UpdatePost(id int, title string, content string) {
	db.posts[id] = models.NewPost(id, title, content)
}

// GetCountPosts get count of all posts
func (db *DB) GetCountPosts() int {
	return len(db.posts)
}

// GetAllPosts get all posts as a map
func (db *DB) GetAllPosts() map[int]models.Post {
	return db.posts
}

// GetPostByID get post by id
func (db *DB) GetPostByID(id int) models.Post {
	return db.posts[id]
}

// Init get db
func Init() DB {
	return DB{
		posts: map[int]models.Post{},
	}
}
