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

	"github.com/gin-gonic/gin"
	"github.com/rawleyfowler/rawleydotxyz/repos"
)

type MusicController struct {
	Repository *repos.MusicRepo
}

func NewMusicController(r *repos.MusicRepo) *MusicController {
	c := new(MusicController)
	c.Repository = r
	return c
}

func RegisterMusicGroup(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "wip.tmpl", &gin.H{})
	})
}

func (mc *MusicController) IndexMusicPage(c *gin.Context) {
	music, err := mc.Repository.GetAllSongs()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "internal_server_error.tmpl", &gin.H{})
		return
	}
	c.HTML(http.StatusOK, "tunes.tmpl", music)
}
