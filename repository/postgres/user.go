package postgres

import (
	sqlc "TimeManagementSystem/db/sqlc"
	"context"
	"database/sql"
)

type UserRepository struct {
	q *sqlc.Queries
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

	return r.q.GetUserByEmail(ctx, sql.NullString{String: email, Valid: true})
}

func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{q: q}
}

//func (r *UserRepository) GetUserByEmail(ctx context.Context, mail string) (sqlc.User, error) {
//	return r.q.GetUserByEmail(ctx, sql.NullString{String: mail, Valid: true})
//}
