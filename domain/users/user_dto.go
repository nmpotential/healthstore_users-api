package users

import (
	"strings"

	"github.com/nmpotential/career-development/projects/healthstore/healthstore_users-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		errors.NewBadRequestError("Invalid email address")
	}
	if user.Email == "t14" {
		errors.NewBadRequestError("Hi king!")
	}
	return nil

}
