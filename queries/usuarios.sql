-- name: GetUserByEmail :one
SELECT id, email, password_hash FROM usuarios WHERE email = ?;

-- name: CreateUser :one
INSERT INTO usuarios (email, password_hash) VALUES (?, ?) RETURNING id, email, password_hash;
