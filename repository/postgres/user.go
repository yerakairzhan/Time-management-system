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
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) GetUserByEmail(email string) (sqlc.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) CreateUser(ctx context.Context, mail string, hashedPassword string) (sqlc.User, error) {
	arg := sqlc.CreateUserParams{
		Email:          sql.NullString{String: mail, Valid: true},
		HashedPassword: sql.NullString{String: hashedPassword, Valid: true},
	}
	return r.q.CreateUser(ctx, arg)
}

//func (r *UserRepository) GetUserByEmail(ctx context.Context, mail string) (sqlc.User, error) {
//	return r.q.GetUserByEmail(ctx, sql.NullString{String: mail, Valid: true})
//}
