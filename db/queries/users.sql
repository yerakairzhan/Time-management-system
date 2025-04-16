-- name: CreateUser :one
INSERT INTO users (email, hashed_password)
VALUES ($1, $2)
    RETURNING id;

-- name: GetUserByEmail :one
select * from users where email = $1;