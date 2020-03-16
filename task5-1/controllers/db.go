package controllers

import (
	"../models"
	"database/sql"
	"fmt"
	"log"
)

const (
	databaseName = "blog"
	tableName    = "posts"
)

func getPosts(db *sql.DB) ([]models.Post, error) {
	res := make([]models.Post, 0, 1)

	rows, err := db.Query(`select * from ` + databaseName + `.` + tableName)
	if err != nil {
		return res, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	for rows.Next() {
		post := models.Post{}

		if err := rows.Scan(&post.ID, &post.Title, &post.Content); err != nil {
			log.Println(err)
			continue
		}

		res = append(res, post)
	}

	return res, nil
}

func getPost(db *sql.DB, id string) (models.Post, error) {
	row := db.QueryRow(fmt.Sprintf(`select * from `+databaseName+`.`+tableName+` WHERE id = %v`, id))

	post := models.Post{}
	if err := row.Scan(&post.ID, &post.Title, &post.Content); err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func addPost(db *sql.DB, post models.Post) error {
	_, err := db.Exec(`INSERT into `+databaseName+`.`+tableName+` (id,title,content) values (?,?,?);`,
		post.ID, post.Title, post.Content)

	return err
}

func editPost(db *sql.DB, post *models.Post, id string) error {
	query := fmt.Sprintf(`UPDATE `+databaseName+`.`+tableName+` SET title="%s", content="%s"  where id=?;`,
		post.Title, post.Content)
	_, err := db.Exec(query, id)

	return err
}

func deletePost(db *sql.DB, id string) error {
	_, err := db.Exec(`DELETE FROM `+databaseName+`.`+tableName+` where id=?;`, id)

	return err
}
