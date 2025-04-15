package postgres

import (
	db "TimeManagementSystem/db/sqlc"
	"context"
	"database/sql"
	"errors"
)

func (r *UserRepository) GetUser(email, hashedPassword string) (db.User, error) {
	user, err := r.q.GetUserByEmail(context.Background(), sql.NullString{String: email, Valid: true})
	if err != nil {
		return db.User{}, err
	}

	if user.HashedPassword.String != hashedPassword {
		return db.User{}, errors.New("invalid password")
	}

	return user, nil
}
