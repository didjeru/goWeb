package db

import (
	"../config"
	"../models"
	"database/sql"
	"log"
)

// DB database of project
type DB struct {
	sql          *sql.DB
	posts        map[int]models.Post
	databaseName string
	tableName    string
}

// AddNewPost add a new post
func (db *DB) AddNewPost(title string, content string) {
	_, err := db.sql.Exec("INSERT INTO "+db.databaseName+"."+db.tableName+" (title,content) VALUES (?,?);", title, content)
	if err != nil {
		log.Println(err)
	}
}

// IsPostExists Check if post is exist
func (db *DB) IsPostExists(id int) bool {
	var isExists bool
	err := db.sql.QueryRow("SELECT EXISTS(SELECT id FROM "+db.tableName+" WHERE id = ?)", id).Scan(&isExists)
	if err != nil {
		log.Println(err)
	}
	return isExists
}

// UpdatePost update an exists post
func (db *DB) UpdatePost(id int, title string, content string) {
	_, err := db.sql.Exec("UPDATE "+db.tableName+" SET title = ?, content = ? WHERE id = ?", title, content, id)
	if err != nil {
		log.Println(err)
	}
}

// GetCountPosts get count of all posts
func (db *DB) GetCountPosts() int {
	var count int
	err := db.sql.QueryRow("SELECT count(*) as count FROM " + db.tableName).Scan(&count)
	if err != nil {
		log.Println(err)
	}
	return count
}

// GetAllPosts get all posts as a map
func (db *DB) GetAllPosts() map[int]models.Post {
	var (
		id      int
		title   string
		content string
	)

	rows, err := db.sql.Query("SELECT id, title, content FROM " + db.tableName)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	posts := make(map[int]models.Post)

	for rows.Next() {
		err := rows.Scan(&id, &title, &content)
		if err != nil {
			log.Println(err)
		}
		posts[id] = models.NewPost(id, title, content)
	}

	return posts
}

// GetPostByID get post by id
func (db *DB) GetPostByID(id int) models.Post {
	var title string
	var content string
	err := db.sql.QueryRow("SELECT title, content FROM "+db.tableName+" WHERE id = ?", id).Scan(&title, &content)
	if err != nil {
		log.Println(err)
	}
	return models.NewPost(id, title, content)
}

// Init get db
func Init(conf config.DatabaseConfig) DB {
	DSN := conf.User + ":" + conf.Password + "@tcp(" + conf.Host + ":" + conf.Port +
		")/blog?charset=utf8"
	db, err := sql.Open("mysql", DSN)

	if err != nil {
		log.Println(err)
	}

	return DB{
		sql:          db,
		posts:        map[int]models.Post{},
		databaseName: conf.Base,
		tableName:    conf.Table,
	}
}
