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
}

func RenderBlogPage(c *gin.Context) {
	c.HTML(http.StatusOK, "blog.tmpl", *GetAllBlogPosts())
}

func RenderIndividualBlogPost(c *gin.Context) {
	c.HTML(http.StatusOK, "blog_post.tmpl", GetBlogPostById(c.Param("url")))
}

func CreateBlogPost(c *gin.Context) {
	reqKey, err := c.Request.Cookie("APK")
	if err != nil {
		c.Status(406)
		return
	}
	if reqKey.Value != apiKey {
		c.Status(403)
		return
	}
	raw, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(406)
		return
	}
	var data models.BlogPost
	// Write the data from the request to the data variable
	if json.Unmarshal([]byte(raw), &data) != nil {
		c.Status(406)
		return
	}
	// Create the blog post in the database
	db.Create(&data)
}

func GetAllBlogPosts() *[]models.BlogPost {
	var posts []models.BlogPost
	db.Find(&posts)
	return &posts
}

func GetBlogPostById(id string) *models.BlogPost {
	var post models.BlogPost
	db.Where(&models.BlogPost{Url: id}).First(&post)
	return &post
}
