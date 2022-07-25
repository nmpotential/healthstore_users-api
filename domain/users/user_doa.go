package users

// dao=data access object
// doa file: will include the entire logic to persist/retrieve user from a given DB
// doa: is the only point in our app where we interact with the DB

import (
	"fmt"

	"github.com/nmpotential/career-development/projects/healthstore/healthstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

/// This method saves a given user
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.Id))
	}
	usersDB[user.Id] = user
	return nil
}
