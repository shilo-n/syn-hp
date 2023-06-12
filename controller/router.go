package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartServer() *gin.Engine {
	r := gin.Default()
	r.Static("/css", "static/css")
	r.Static("/js", "static/js")
	r.Static("/img", "static/img")
	r.LoadHTMLGlob("template/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/profile", func(c *gin.Context) {
		c.HTML(http.StatusOK, "profile.html", nil)
	})
	r.GET("/movie", func(c *gin.Context) {
		c.HTML(http.StatusOK, "movie.html", nil)
	})
	r.GET("/link", func(c *gin.Context) {
		c.HTML(http.StatusOK, "link.html", nil)
	})
	r.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", nil)
	})

	return r
}
