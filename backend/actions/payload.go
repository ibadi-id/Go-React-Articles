package actions

// package ini hanya opsional untuk control payload dari request dan response
// tidak terdapat dalam kebutuhan test hanya sebagai tambahan fungsi

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/ibadi-id/article-test/models"
)

//
// Request dan Response payloads untuk REST api.
//

type ArticleRequest struct {
	Article *models.Posts
	// ProtectedID string `json:"id"` // override 'id' json to have more control
}

func (a *ArticleRequest) Bind(r *http.Request) error {
	// Antisipasi untuk nil pointer dereference.
	if a.Article == nil {
		return errors.New("missing required Article fields")
	}
	// // Hanya contoh untuk edit data payload request
	// a.ProtectedID = ""  // unset the protected ID
	// a.Article.Title = strings.ToLower(a.Article.Title) // sebagai contoh title di lower case
	return nil
}

// ArticleResponse adalah type response payload untuk Model Artikel.
type ArticleResponse struct {
	*models.Posts
	// Contoh Penambahan data Pada response
	// Elapsed int64 `json:"elapsed"`
}

func NewArticleResponse(article *models.Posts) *ArticleResponse {
	resp := &ArticleResponse{article}
	return resp
}

func (rd *ArticleResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Hanya contoh apabila ingin menambah atau mengubah data sebelum Respon ke User
	// rd.Elapsed = 10
	return nil
}

func NewArticleListResponse(articles []*models.Posts) []render.Renderer {
	list := []render.Renderer{}
	for _, article := range articles {
		list = append(list, NewArticleResponse(article))
	}
	return list
}

//--
// Error response payloads & renderers
//--

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
