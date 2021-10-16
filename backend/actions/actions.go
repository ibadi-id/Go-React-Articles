package actions

// Package action untuk respon request dari client : CRUD Article

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/ibadi-id/article-test/models"
)

// ListArticles adalah fungsi untuk menampilkan semua Article
func ListArticles(w http.ResponseWriter, r *http.Request) {
	var err error
	var articles []*models.Posts

	articles, err = dbListArticle()
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}


	if err := render.RenderList(w, r, NewArticleListResponse(articles)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

// ListArticlesWithPage adalah fungsi untuk menampilkan article dengan pagination
func ListArticlesWithPage(w http.ResponseWriter, r *http.Request) {
	var err error
	var lim_int int
	var off_int int
	var articles []*models.Posts

	limit := chi.URLParam(r, "limit")
	offset := chi.URLParam(r, "offset")

	lim_int, err = strconv.Atoi(limit)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	off_int, err = strconv.Atoi(offset)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	articles, err = dbListArticlePage(&lim_int, &off_int)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}


	if err := render.RenderList(w, r, NewArticleListResponse(articles)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

// ArticleCtx middleware untuk load context "article" pada request
func ArticleCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var article *models.Posts
		var err error

		articleID := chi.URLParam(r, "articleID")

		id_int, err := strconv.Atoi(articleID)

		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		article, err = dbGetArticle(id_int)
		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "article", article)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// CreateArticle adalah fungsi untuk persists article baru dan mengembalikannya
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	data := &ArticleRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	article := data.Article
	id, _ := dbNewArticle(article)

	data.Article.Id = int(id)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewArticleResponse(article))
}

// GetArticle fungsi untuk mengembalikan spesifik article (default by id, bisa dikembangkan by slug)
func GetArticle(w http.ResponseWriter, r *http.Request) {

	article := r.Context().Value("article").(*models.Posts)

	if err := render.Render(w, r, NewArticleResponse(article)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

// UpdateArticle fungsi untuk updates Article dan mengembalikannya
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	article := r.Context().Value("article").(*models.Posts)

	data := &ArticleRequest{Article: article}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	article = data.Article
	dbUpdateArticle(article.Id, article)

	render.Render(w, r, NewArticleResponse(article))
}

// DeleteArticle adalah fungsi untuk hapus article by id
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	var err error
	article := r.Context().Value("article").(*models.Posts)

	dbRemoveArticle(article.Id)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, NewArticleResponse(article))
}