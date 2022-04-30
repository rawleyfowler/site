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
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/rawleyfowler/rawleydotxyz/models"
	"github.com/rawleyfowler/rawleydotxyz/utils"
	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(dsnPath string) *AdminRepo {
	c := AdminRepo{}
	var err error
	c.DB, err = utils.CreateDatabase(utils.LoadDSN(dsnPath))
	if err != nil {
		panic(err)
	}
	c.DB.AutoMigrate(&models.Administrator{})
	return &c
}

func (ar *AdminRepo) CreateAdmin(a *models.Administrator) error {
	if a.Password == "" ||
		a.Username == "" {
		return errors.New("Administrators must have a valid username and password")
	}
	if len(a.Password) < 8 {
		return errors.New("Passwords must have a length of atleast 8")
	}
	a.Token = uuid.New().String()
	a.Password = fmt.Sprintf("%x", string(sha256.New().Sum([]byte(a.Password))))
	err := ar.DB.Table("administrators").Save(a).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *AdminRepo) GetAdminByToken(k string) (*models.Administrator, error) {
	if k == "" {
		return nil, errors.New("Session key cannot be blank")
	}
	if len(k) != 36 {
		return nil, errors.New("Valid sessions keys are always 36 byte utf-8 strings")
	}
	var a *models.Administrator
	err := ar.DB.Table("administrators").Where(&models.Administrator{Token: k}).Scan(a).Error
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (ar *AdminRepo) GetAdminByCredentials(uname, pword string) (*models.Administrator, error) {
	if uname == "" ||
		pword == "" {
		return nil, errors.New("Username or password mustn't be empty")
	}
	if len(pword) < 8 {
		return nil, errors.New("Password length must be no less than 8")
	}
	pword = fmt.Sprintf("%x", string(sha256.New().Sum([]byte(pword))))
	var a models.Administrator
	err := ar.DB.Table("administrators").Where(&models.Administrator{Username: uname, Password: pword}).First(&a).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}
