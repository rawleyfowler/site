package repos

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
	"errors"

	"gorm.io/gorm"

	"github.com/rawleyfowler/rawleydotxyz/models"
	"github.com/rawleyfowler/rawleydotxyz/utils"
)

type MusicRepo struct {
	DB *gorm.DB
}

func NewMusicRepo(dsnPath string) *MusicRepo {
	mr := new(MusicRepo)
	var err error
	mr.DB, err = utils.CreateDatabase(utils.LoadDSN(dsnPath))
	if err != nil {
		panic(err)
	}
	return mr
}

func (mr *MusicRepo) GetAllSongs() (*[]models.Music, error) {
	songs := new([]models.Music)
	err := mr.DB.Table("music").Order("name desc").Scan(songs).Error
	if err != nil {
		return nil, errors.New("Could not load music from database")
	}
	return songs, nil
}

func (mr *MusicRepo) CreateMusic(m *models.Music) error {
	if m.Name == "" || m.Url == "" {
		return errors.New("Cannot add music with empty or partially filled values")
	}
	if err := mr.DB.Table("music").Save(m).Error; err != nil {
		return err
	}
	return nil
}
