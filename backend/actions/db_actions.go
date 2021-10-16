package actions

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ibadi-id/article-test/database"
	"github.com/ibadi-id/article-test/models"
)

// dbNewArticle adalah fungsi untuk menambah artcle ke database
func dbNewArticle(article *models.Posts) (int64, error){

	stmt, err := database.Db.Prepare("INSERT INTO posts (title, content, category, status) VALUES(?,?,?,?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	
	res, err :=  stmt.Exec(&article.Title, &article.Content, &article.Category, &article.Status)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
	
}

// dbListArticle adalah fungsi untuk semua artikel ke database
func dbListArticle() ([]*models.Posts, error) {

	var articles []*models.Posts

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err:= database.Db.QueryContext(ctx, "SELECT * FROM posts")
	if err != nil {
		fmt.Println("error", err)
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Posts
		err := rows.Scan(
			&i.Id,
			&i.Title,
			&i.Content,
			&i.Category,
			&i.CreatedDate,
			&i.UpdatedDate,
			&i.Status,
		)
		if err != nil {
			fmt.Println("error", err)
		}
		articles = append(articles, &i)
	}

	return articles, nil
}

// dbListArticlePage adalah fungsi untuk get list artikel dengan parameter (limit, offset) ke database
func dbListArticlePage(limit *int, offset *int) ([]*models.Posts, error) {

	var articles []*models.Posts

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err:= database.Db.QueryContext(ctx, "SELECT * FROM posts  WHERE status = 'Publish' LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		fmt.Println("error", err)
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Posts
		err := rows.Scan(
			&i.Id,
			&i.Title,
			&i.Content,
			&i.Category,
			&i.CreatedDate,
			&i.UpdatedDate,
			&i.Status,
		)
		if err != nil {
			fmt.Println("error", err)
		}
		articles = append(articles, &i)
	}

	return articles, nil
}

// dbGetArticle adalah fungsi untuk get artikel by id ke database
func dbGetArticle(id int) (*models.Posts, error) {
	var article models.Posts
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM posts WHERE id = ?"
	

	err := database.Db.QueryRowContext(ctx, query, id).Scan( 
		&article.Id, &article.Title, &article.Content, &article.Category,
		&article.CreatedDate, &article.UpdatedDate, &article.Status,
	)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal database error")
	}
	return &article, nil
}

// dbUpdateArticle adalah fungsi untuk update artikel by id ke database
func dbUpdateArticle(id int, article *models.Posts) (int64, error) {
	stmt, err := database.Db.Prepare("UPDATE posts SET title = ?, content = ?, category = ?, status = ? WHERE id = ?")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	
	res, err :=  stmt.Exec(&article.Title, &article.Content, &article.Category, &article.Status, id)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
	
}

// dbUpdateArticle adalah fungsi untuk hapus artikel by id dari database
func dbRemoveArticle(id int) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	fmt.Println("id", id)

	query := "DELETE FROM posts WHERE id = ?"
	

	err := database.Db.QueryRowContext(ctx, query, id)

	if err != nil {
		fmt.Println(err)
	}
}
