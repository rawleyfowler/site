package repos

import (
	"errors"

	"gitlab.com/rawleyifowler/site-rework/models"
	"gitlab.com/rawleyifowler/site-rework/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CommentRepo struct {
	DB          *gorm.DB
	Initialized bool
}

func NewCommentRepo(dsnPath string) (*CommentRepo, error) {
	c := CommentRepo{}
	err := c.Initialize(dsnPath)
	if err != nil {
		return &c, err
	}
	return &c, nil
}

func (c *CommentRepo) Initialize(path string) error {
	dsn := utils.LoadDSN(path)
	var err error
	c.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	c.Initialized = true
	return nil
}

func (c *CommentRepo) GetCommentById(id uint) (*models.Comment, error) {
	if id == 0 {
		return nil, nil
	}
	var comm models.Comment
	c.DB.Where(&models.Comment{Id: id}).First(&comm)
	return &comm, nil
}

func (c *CommentRepo) CreateComment(comm *models.Comment) error {
	if comm == nil {
		return errors.New("nil comment")
	}
	err := c.DB.Create(comm).Error
	if err != nil {
		return err
	}
	// nil return means no errors, yay!
	return nil
}

func (c *CommentRepo) GetCommentsByAssociatedPost(title string) (*[]models.Comment, error) {
	if len(title) < 1 {
		return nil, errors.New("empty associated post title")
	}
	var comms []models.Comment
	err := c.DB.Where(&models.Comment{AssociatedPost: title}).Scan(&comms).Error
	if err != nil {
		return nil, err
	}
	return &comms, nil
}

func (c *CommentRepo) GetCommentsByAuthor(auth string) (*[]models.Comment, error) {
	if len(auth) < 1 {
		return nil, errors.New("empty author name")
	}
	var comms []models.Comment
	err := c.DB.Where(&models.Comment{Author: auth}).Scan(&comms).Error
	if err != nil {
		return nil, err
	}
	return &comms, nil
}
