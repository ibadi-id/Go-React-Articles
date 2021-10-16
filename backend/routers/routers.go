package routers

// Package ini untuk mengatur routing dari server

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"
	"github.com/ibadi-id/article-test/actions"
	"github.com/ibadi-id/article-test/helper"
)

// code ini untuk generate dokumentasi, default false
// apabila di set "true" maka server hanya akan generate Dokumentasi
var routes = flag.Bool("routes", false, "Generate router documentation")

func ApiRouter() {
	flag.Parse()

	// code ini opsional, fungsinya jika terjadi error pada code,
	// maka akan menginformasikan nama file dan baris code
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// routes menggunakan pacakge go-chi
	r := chi.NewRouter()

	// middleware yang digunakan pada routes bersifat opsional sesuai kebutuhan
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		Debug:            false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Daftar routes untuk API REST public
	r.Route("/article", func(r chi.Router) {
		r.Get("/", actions.ListArticles)
		r.Get("/{limit}/{offset}", actions.ListArticlesWithPage)
		r.Post("/", actions.CreateArticle)       // POST /articles
		// r.Get("/search", actions.SearchArticles) // GET /articles/search

		r.Route("/{articleID}", func(r chi.Router) {
			r.Use(actions.ArticleCtx)            // Load the *Article on the request context
			r.Get("/", actions.GetArticle)       // GET /articles/123
			r.Put("/", actions.UpdateArticle)    // PUT /articles/123
			r.Delete("/", actions.DeleteArticle) // DELETE /articles/123
		})
	})


	// Code ini Untuk generate dokumentasi REST API dengan perintah "go run . -routes"
	if *routes {
		a := docgen.JSONRoutesDoc(r)
		b := docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "./docs",
			ForceRelativeLinks: false,
			Intro:       "Welcome to the article-test/rest generated docs.",
		})
		_ = ioutil.WriteFile("./docs/docs.md", []byte(b), 0644)
		_ = ioutil.WriteFile("./docs/routers.json", []byte(a), 0644)
		return
	}

	// Code ini untuk menampilkan pesan pada terminal
	color.Green("Listening on localhost%s\n", ":3333")
	helper.LineSeparator()

	http.ListenAndServe(":3333", r)
}
