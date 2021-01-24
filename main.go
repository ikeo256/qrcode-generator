package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	
	r.LoadHTMLGlob("templates/*.html") // HTMLファイルがあるディレクトリを指定

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{}) // サイトのトップにアクセスがきたらindex.htmlにルーティング
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}