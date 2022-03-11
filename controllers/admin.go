package controllers

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/rawleyifowler/site-rework/models"
	"gitlab.com/rawleyifowler/site-rework/utils"
)

var (
	key        string            = utils.LoadApiKey("api_key.pem")
	admin_hash [32]byte          = utils.LoadAdminHash("admin_hash.pem")
	att        map[string]uint32 = make(map[string]uint32)
)

// Ensure is logged in
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		var status uint = http.StatusOK
		k, err := c.Cookie("admin_key")
		if err != nil {
			status = http.StatusUnauthorized
		}
		fmt.Println(k)
		if k == key && status == http.StatusOK {
			c.Next()
		} else {
			c.HTML(int(status), "forbidden.tmpl", models.Page{Title: " | forbidden"})
			c.Abort()
			return
		}
	}
}

func RegisterAdminGroup(r *gin.RouterGroup) {
	r.GET("/", utils.ServePage("login.tmpl"))
	r.POST("/login", HandleLogin)
	r.GET("/post", AdminOnly(), utils.ServePage("create_post.tmpl"))
	r.POST("/post", AdminOnly(), CreatePost)
}

func HandleLogin(c *gin.Context) {
	if att[c.ClientIP()] > 5 {
		c.HTML(http.StatusUnauthorized, "forbidden.tmpl", models.Page{Title: " | forbidden"})
		return
	}
	err := c.Request.ParseForm()
	if err != nil {
		c.HTML(http.StatusUnauthorized, "forbidden.tmpl", models.Page{Title: " | forbidden"})
		return
	}
	f := c.Request.Form
	for _, v := range []string{"username", "password"} {
		if !f.Has(v) {
			c.AbortWithStatus(http.StatusNotAcceptable)
			return
		}
	}
	str := f.Get("username") + f.Get("password")
	s := sha256.Sum256([]byte(str))
	var success bool
	if s == admin_hash {
		c.SetCookie("admin_key", key, 3600, "/", "rawley.xyz", false, true)
		success = true
	} else {
		att[c.ClientIP()]++
		success = false
	}
	c.HTML(http.StatusOK, "admin_success_redirect.tmpl", struct {
		Success bool
		Title   string
	}{Success: success, Title: "login"})
}

func CreatePost(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	f := c.Request.Form
	for _, v := range []string{"title", "content", "url"} {
		if !f.Has(v) {
			c.AbortWithStatus(http.StatusNotAcceptable)
			return
		}
	}
	b := AddBlogPost(&models.BlogPost{
		Title:   f.Get("title"),
		Content: f.Get("content"),
		Url:     f.Get("url"),
	})
	c.HTML(http.StatusOK, "post_success.tmpl", struct {
		Url     string
		Success bool
		Title   string
	}{Url: f.Get("url"), Success: b, Title: " | post attempt"})
}
