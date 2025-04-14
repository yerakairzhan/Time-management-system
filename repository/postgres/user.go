package postgres

import (
	sqlc "TimeManagementSystem/db/sqlc"
	"context"
)

type UserRepository struct {
	q *sqlc.Queries
}

func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) CreateUser(ctx context.Context, mail string, hashedPassword string) (sqlc.User, error) {
	arg := sqlc.CreateUserParams{
		Email:          mail,
		HashedPassword: hashedPassword,
	}
	return r.q.CreateUser(ctx, arg)
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, mail string) (sqlc.User, error) {
	return r.q.GetUserByEmail(ctx, mail)
}
