package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/ibadi-id/article-test/database"
	"github.com/ibadi-id/article-test/helper"
	"github.com/ibadi-id/article-test/routers"
)

func main() {
	// perintah untuk pesan awal pada terminal
	helper.StartMessage("Memulai Server Artikel Sharing Vision")

	// perintah untuk membuat koneksi dengan database
	database.InitDB()
	defer database.Db.Close()
	
	// perintah untuk menjalankan routers
	routers.ApiRouter()
}