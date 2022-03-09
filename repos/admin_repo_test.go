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
	"testing"

	"gitlab.com/rawleyifowler/site-rework/models"
)

var test_admin_repo = NewAdminRepo("dsn")

func TestCreateAdmin(t *testing.T) {
	err := test_admin_repo.CreateAdmin(&models.Administrator{})
	if err == nil {
		t.Fatalf("Expected: You cannot create an admin with empty credentials")
	}
	err = test_admin_repo.CreateAdmin(&models.Administrator{Username: "", Password: "12345678"})
	if err == nil {
		t.Fatalf("Expected: You cannot create an admin with empty credentials")
	}
	err = test_admin_repo.CreateAdmin(&models.Administrator{Username: "12345678", Password: ""})
	if err == nil {
		t.Fatalf("Expected: You cannot create an admin with empty credentials")
	}
	err = test_admin_repo.CreateAdmin(&models.Administrator{Username: "valid_username", Password: "invalid"})
	if err == nil {
		t.Fatalf("Expected: You cannot create an admin with a password less than 8 characters")
	}
}

func TestGetAdminByToken(t *testing.T) {
	_, err := test_admin_repo.GetAdminByToken("")
	if err == nil {
		t.Fatalf("Expected: Error should be thrown when an empty session key is provided")
	}
}

func TestGetAdminByCredentials(t *testing.T) {
	_, err := test_admin_repo.GetAdminByCredentials("", "")
	if err == nil {
		t.Fatalf("Expected: Error should be thrown if empty credentials are given")
	}
	_, err = test_admin_repo.GetAdminByCredentials("1234", "")
	if err == nil {
		t.Fatalf("Expected: Error should be thrown if empty credentials are given")
	}
	_, err = test_admin_repo.GetAdminByCredentials("", "123456789")
	if err == nil {
		t.Fatalf("Expected: Error should be thrown if empty credentials are given")
	}
}
