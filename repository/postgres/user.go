package postgres

import (
	sqlc "TimeManagementSystem/db/sqlc"
	"context"
	"errors"
)

type UserRepository struct {
	q *sqlc.Queries
}

func (r *UserRepository) GetUser(email, hashedPassword string) (sqlc.User, error) {
	user, err := r.q.GetUserByEmail(context.Background(), email)
	if err != nil {
		return sqlc.User{}, err
	}

	if user.HashedPassword != hashedPassword {
		return sqlc.User{}, errors.New("invalid password")
	}

	return user, nil
}

func (r *UserRepository) Create(user sqlc.User) (int, error) {
	ctx := context.Background()

	arg := sqlc.CreateUserParams{
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
	}

	id, err := r.q.CreateUser(ctx, arg)
	return int(id), err
}

func (r *UserRepository) GetUserByEmail(email string) (sqlc.User, error) {
	ctx := context.Background()

	return r.q.GetUserByEmail(ctx, email)
}

func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{q: q}
}

//func (r *UserRepository) GetUserByEmail(ctx context.Context, mail string) (sqlc.User, error) {
//	return r.q.GetUserByEmail(ctx, sql.NullString{String: mail, Valid: true})
//}
