package handler

import (
	"golang.org/x/crypto/bcrypt"

	db "TimeManagementSystem/db/sqlc"
)

type CreateUserMatcher struct {
	ExpectedEmail     string
	PlaintextPassword string
}

func (m CreateUserMatcher) Matches(x interface{}) bool {
	user, ok := x.(db.User)
	if !ok {
		return false
	}

	if user.Email != m.ExpectedEmail {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(m.PlaintextPassword))
	return err == nil
}

func (m CreateUserMatcher) String() string {
	return "matches user with email=" + m.ExpectedEmail + " and correct password hash"
}
