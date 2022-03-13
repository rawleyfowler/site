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

	"gitlab.com/rawleyifowler/site-rework/models"
	"gitlab.com/rawleyifowler/site-rework/utils"
	"gorm.io/gorm"
)

type BlogRepo struct {
	DB          *gorm.DB
	initialized bool
}

func NewBlogRepo(dsnPath string) *BlogRepo {
	b := BlogRepo{}
	var err error
	b.DB, err = utils.CreateDatabase(utils.LoadDSN(dsnPath))
	if err != nil {
		panic(err)
	}
	b.DB.AutoMigrate(&models.BlogPost{})
	return &b
}

func (b *BlogRepo) CreateBlogPost(p *models.BlogPost) error {
	if p == nil {
		return errors.New("Blog post must not be nil")
	}
	if p.Title == "" ||
		p.Content == "" ||
		p.Url == "" {
		return errors.New("Blog posts must be fully formed")
	}
	err := b.DB.Table("blog_posts").Create(p).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *BlogRepo) GetBlogByUrl(url string) (*models.BlogPost, error) {
	if url == "" {
		return nil, errors.New("Cannot query by empty value")
	}
	var post models.BlogPost
	err := b.DB.Table("blog_posts").Where(&models.BlogPost{Url: url}).Scan(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (b *BlogRepo) GetAllBlogPosts() (*[]models.BlogPost, error) {
	posts := new([]models.BlogPost)
	err := b.DB.Table("blog_posts").Order("date desc").Scan(posts).Error
	if err != nil {
		return nil, errors.New("Could not load all posts from database")
	}
	return posts, nil
}

func (b *BlogRepo) UpdateExistingPost(p *models.BlogPost) error {
	if p == nil {
		return errors.New("Cannot update with nil value for post")
	}
	if p.Url == "" {
		return errors.New("Url must exist to update a blog post record")
	}
	err := b.DB.Table("blog_posts").Where(&models.BlogPost{Url: p.Url}).Save(p).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *BlogRepo) DeleteExistingPost(p *models.BlogPost) error {
	if p == nil {
		return errors.New("Cannot delete with nil value for post")
	}
	if p.Url == "" {
		return errors.New("Url must exist to delete a blog post record")
	}
	err := b.DB.Table("blog_posts").Delete(p).Error
	if err != nil {
		return err
	}
	return nil
}
