package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Returns a handler function to serve a page based on the template that is inputted
func ServePage(temp string) gin.HandlerFunc {
	title := strings.Split(temp, ".")[0]
	if title == "index" {
		// Index is the home page so we don't care
		title = ""
	} else {
		// Else title it accordingly
		title = " | " + title
	}
	return func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			temp,
			struct{ Title string }{Title: title},
		)
	}
}
