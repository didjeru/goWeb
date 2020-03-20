package db

import (
	"../models"
	"sync"
)

// DB database of project
type DB struct {
	sync.Mutex
	posts map[int]models.Post
}

// AddNewPost add a new post
func (db *DB) AddNewPost(title string, content string) {
	id := db.GetCountPosts()
	db.Lock()
	db.posts[id] = models.NewPost(id, title, content)
	db.Lock()
}

// IsPostExists Check if post is exist
func (db *DB) IsPostExists(id int) bool {
	db.Lock()
	_, ok := db.posts[id]
	db.Unlock()
	return ok
}

// UpdatePost update an exists post
func (db *DB) UpdatePost(id int, title string, content string) {
	db.Lock()
	db.posts[id] = models.NewPost(id, title, content)
	db.Unlock()
}

// GetCountPosts get count of all posts
func (db *DB) GetCountPosts() int {
	db.Lock()
	defer db.Unlock()
	return len(db.posts)
}

// GetAllPosts get all posts as a map
func (db *DB) GetAllPosts() map[int]models.Post {
	db.Lock()
	defer db.Unlock()
	return db.posts
}

// GetPostByID get post by id
func (db *DB) GetPostByID(id int) models.Post {
	db.Lock()
	defer db.Unlock()
	return db.posts[id]
}

// Init get db
func Init() DB {
	return DB{
		posts: map[int]models.Post{},
	}
}
