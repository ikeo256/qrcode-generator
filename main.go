package main

import (
	"net/http"
	"html/template"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*.html") // HTMLファイルがあるディレクトリを指定

	r.GET("/", func(ctx *gin.Context) {
		html := template.Must(template.ParseFiles("templates/navbar.html",
		"templates/index.html"))

        r.SetHTMLTemplate(html)
		ctx.HTML(http.StatusOK, "navbar.html", gin.H{}) // サイトのトップにアクセスがきたらindex.htmlにルーティング
	})

	r.GET("/settings", func(ctx *gin.Context) {
		html := template.Must(template.ParseFiles("templates/navbar.html",
		"templates/settings.html"))

		r.SetHTMLTemplate(html)
		ctx.HTML(http.StatusOK, "navbar.html", gin.H{})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}