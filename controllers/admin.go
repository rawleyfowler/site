package controllers

import (
	"crypto/sha256"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/rawleyifowler/site-rework/models"
	"gitlab.com/rawleyifowler/site-rework/utils"
)

var (
	key        string            = utils.LoadApiKey("api_key.pem")
	admin_hash string            = utils.LoadAdminHash("admin_hash.pem")
	att        map[string]uint32 = make(map[string]uint32)
)

// Ensure is logged in
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		k, err := c.Cookie("admin_key")
		if err != nil {
			c.Redirect(http.StatusPermanentRedirect, "/admin/")
		}
		if k == key {
			c.Next()
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
		c.Redirect(http.StatusPermanentRedirect, "/admin/")
	}
	err := c.Request.ParseForm()
	if err != nil {
		c.Redirect(http.StatusPermanentRedirect, "/admin/")
	}
	f := c.Request.Form
	for _, v := range []string{"username", "password"} {
		if !f.Has(v) {
			c.AbortWithStatus(http.StatusNotAcceptable)
		}
	}
	str := f.Get("username") + f.Get("password")
	s := sha256.New()
	s.Sum([]byte(str))
	var h []byte
	s.Write(h)
	if string(h) == admin_hash {
		c.SetCookie("admin_key", key, 1, "/", "rawley.xyz", true, true)
		c.Redirect(http.StatusTemporaryRedirect, "/admin/post")
	} else {
		att[c.ClientIP()]++
		c.Redirect(http.StatusTemporaryRedirect, "/admin/")
	}
}

func CreatePost(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	f := c.Request.Form
	for _, v := range []string{"title", "content", "url"} {
		if !f.Has(v) {
			c.AbortWithStatus(http.StatusNotAcceptable)
		}
	}
	b := AddBlogPost(&models.BlogPost{
		Title:   f.Get("title"),
		Content: f.Get("content"),
		Url:     f.Get("url"),
	})
	if b {
		c.Redirect(http.StatusTemporaryRedirect, "/blog/post/"+f.Get("url"))
	} else {
		c.Redirect(http.StatusTemporaryRedirect, "/admin/post")
	}
}
