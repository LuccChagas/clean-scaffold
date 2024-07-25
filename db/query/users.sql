
-- name: CreateUser :one
INSERT INTO users (id, age, name, birth_date, status)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetAllUsers :many
SELECT id, age, name, birth_date, status FROM users;


-- name: GetUsersByID :one
SELECT id, age, name, birth_date, status FROM users WHERE id = $1;


-- name: UpdateUser :one
UPDATE users
SET age = $2, name = $3, birth_date = $4, status = $5
WHERE id = $1
RETURNING *;

-- DeleteUser :one
UPDATE users
SET status = FALSE
WHERE id = $1
RETURNING *;
