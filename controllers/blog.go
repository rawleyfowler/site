package controllers

/* Copyright (C) 2022 Rawley Fowler
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
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/rawleyifowler/site-rework/models"
	"gitlab.com/rawleyifowler/site-rework/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	apiKey string
)

func RegisterBlogGroup(r *gin.RouterGroup) {
	// Load dsn and initialize database
	dsn := utils.LoadDSN("dsn")
	apiKey = utils.LoadApiKey("api_key.pem")
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connection failed...")
	}
	utils.PerformMigrations(db)
	r.GET("/", RenderBlogPage)
	r.GET("/post/:url", RenderIndividualBlogPost)
	r.POST("/post", CreateBlogPost)
	r.POST("/post/comment", CreateComment)
}

func RenderBlogPage(c *gin.Context) {
	c.HTML(http.StatusOK, "blog.tmpl", *GetAllBlogPosts())
}

func RenderIndividualBlogPost(c *gin.Context) {
	bp := GetBlogPostById(c.Param("url"))
	if bp == nil {
		c.HTML(http.StatusNotFound, "not_found.tmpl", gin.H{})
	} else {
		c.HTML(http.StatusOK, "blog_post.tmpl", bp)
	}
}

func CreateBlogPost(c *gin.Context) {
	reqKey, err := c.Request.Cookie("APK")
	if err != nil {
		c.Status(http.StatusForbidden)
		return
	}
	if reqKey.Value != apiKey {
		c.Status(http.StatusForbidden)
		return
	}
	raw, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusNotAcceptable)
		return
	}
	var data models.BlogPost
	// Write the data from the request to the data variable
	if json.Unmarshal([]byte(raw), &data) != nil {
		c.Status(http.StatusNotAcceptable)
		return
	}
	// Create the blog post in the database
	if db.Create(&data).Error != nil {
		c.Status(http.StatusNotAcceptable)
	} else {
		c.Status(http.StatusAccepted)
	}
}

func CreateComment(c *gin.Context) {
	if c.Request.ParseForm() != nil {
		c.Status(http.StatusNotAcceptable)
		return
	}
	comment := models.Comment{
		Author:         c.Request.Form.Get("author"),
		Content:        c.Request.Form.Get("content"),
		AssociatedPost: c.Request.Form.Get("url"),
	}
	db.Create(&comment)
}

func GetAllBlogPosts() *[]models.BlogPost {
	var posts []models.BlogPost
	// Select title, date, and url fields from the blog post records an store them in posts.
	// This is so we don't grab the entire blog post when we render them on the overview page. Saves a couple ms.
	db.Model(&models.BlogPost{}).Select("title, date, url").Take(&posts)
	return &posts
}

func GetBlogPostById(id string) *models.BlogPost {
	var post models.BlogPost
	err := db.Where(&models.BlogPost{Url: id}).First(&post).Error
	if err != nil {
		return nil
	}
	// TODO: Figure out gorm joins!
	// gorm joins are not working at all, just going to do another query to make it work for now.
	err = db.Where(&models.Comment{AssociatedPost: id}).Find(&post.Comments).Error
	if err != nil {
		return nil
	}
	return &post
}
