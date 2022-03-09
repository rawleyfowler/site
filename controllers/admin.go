package controllers

/*
Copyright (C) 2022 Rawley Fowler

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.Rawley Fowler, 2022
*/

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/rawleyifowler/site-rework/models"
	"gitlab.com/rawleyifowler/site-rework/repos"
	"gitlab.com/rawleyifowler/site-rework/utils"
)

var (
	api_key string = utils.LoadApiKey("api_key.pem")
)

type AdminController struct {
	Repository        *repos.AdminRepo
	BlogRepository    *repos.BlogRepo
	ActiveLogins      map[string]int64
	LoginAttemptsByIp map[string]uint
}

func NewAdminController(r *repos.AdminRepo) *AdminController {
	var a AdminController
	a.Repository = r
	a.ActiveLogins = map[string]int64{}
	a.LoginAttemptsByIp = map[string]uint{}
	go utils.TimeClearMap(a.ActiveLogins, 850)
	return &a
}

func RegisterAdminGroup(r *gin.RouterGroup) {
	a := NewAdminController(repos.NewAdminRepo("dsn"))
	r.Use(a.UserBannedMiddleware)
	r.GET("/", utils.ServePage("login.tmpl"))
	r.POST("/login", a.AdminLogin)
	r.GET("/post", a.AuthMiddleware, utils.ServePage("create_post.tmpl"))
	r.POST("/post", a.AuthMiddleware, a.CRUDPost)
	r.POST("/user", a.CreateAdmin)
}

func (a *AdminController) AuthMiddleware(c *gin.Context) {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		c.HTML(http.StatusForbidden, "forbidden.tmpl", &gin.H{})
		c.Abort()
		return
	}
	if a.ActiveLogins[cookie.Value] == 0 {
		c.HTML(http.StatusForbidden, "session_revoked.tmpl", &gin.H{})
		c.Abort()
		return
	}
	c.Next()
}

func (a *AdminController) UserBannedMiddleware(c *gin.Context) {
	if a.IpIsBanned(c.ClientIP()) {
		c.HTML(http.StatusForbidden, "forbidden.tmpl", &gin.H{})
		c.Abort()
		return
	}
	c.Next()
}

func (a *AdminController) IpIsBanned(ip string) bool {
	return a.LoginAttemptsByIp[ip] >= 3
}

func (a *AdminController) AdminLogin(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.HTML(http.StatusForbidden, "forbidden.tmpl", &gin.H{})
		a.LoginAttemptsByIp[c.ClientIP()]++
		return
	}
	f := c.Request.Form
	if !f.Has("username") ||
		!f.Has("password") {
		c.HTML(http.StatusNotAcceptable, "admin_success_redirect.tmpl", false)
		a.LoginAttemptsByIp[c.ClientIP()]++
		return
	}
	ad, err := a.Repository.GetAdminByCredentials(f.Get("username"), f.Get("password"))
	if err != nil ||
		ad == nil {
		c.HTML(http.StatusNotAcceptable, "admin_success_redirect.tmpl", false)
		a.LoginAttemptsByIp[c.ClientIP()]++
		return
	}
	a.ActiveLogins[ad.Token] = time.Now().UnixMilli() + (3600 * 1000)
	c.SetCookie("token", ad.Token, 3600*1000, "/", "rawley.xyz", true, true)
	c.HTML(http.StatusAccepted, "admin_success_redirect.tmpl", true)
}

func (a *AdminController) CRUDPost(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.HTML(http.StatusNotAcceptable, "post_success.tmpl", false)
		return
	}
	f := c.Request.Form
	if f.Get("url") == "" ||
		f.Get("op") == "" {
		c.HTML(http.StatusNotAcceptable, "post_success.tmpl", false)
		return
	}
	tempBlogRepo := repos.NewBlogRepo("dsn")
	tempPost := &models.BlogPost{
		Url:     f.Get("url"),
		Content: f.Get("content"),
		Title:   f.Get("title"),
	}
	switch f.Get("op") {
	case "create":
		err = tempBlogRepo.CreateBlogPost(tempPost)
		break
	case "update":
		err = tempBlogRepo.UpdateExistingPost(tempPost)
		break
	case "delete":
		err = tempBlogRepo.DeleteExistingPost(tempPost)
		break
	default:
		c.HTML(http.StatusNotAcceptable, "post_success.tmpl", false)
		return
	}
	if err != nil {
		c.HTML(http.StatusNotAcceptable, "post_success.tmpl", false)
		return
	}
	c.HTML(http.StatusAccepted, "post_success.tmpl", true)
}

func (a *AdminController) CreateAdmin(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.HTML(http.StatusNotAcceptable, "forbidden.tmpl", &gin.H{})
		return
	}
	var at models.Administrator
	f := c.Request.Form
	if !f.Has("password") ||
		!f.Has("username") ||
		!f.Has("api_key") {
		c.HTML(http.StatusBadRequest, "forbidden.tmpl", &gin.H{})
		return
	}
	at.Username = f.Get("username")
	at.Password = f.Get("password")
	if f.Get("api_key") != api_key {
		c.HTML(http.StatusForbidden, "forbidden.tmpl", &gin.H{})
		return
	}
	err = a.Repository.CreateAdmin(&at)
	if err != nil {
		c.HTML(http.StatusForbidden, "forbidden.tmpl", &gin.H{})
		return
	}
	c.HTML(http.StatusAccepted, "post_success.tmpl", true)
}
