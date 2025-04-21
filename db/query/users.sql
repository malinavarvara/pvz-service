-- name: CreateUser :one
INSERT INTO users (
  email,
  password_hash,
  role
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1 LIMIT 1;

-- name: ListUsersByRole :many
SELECT * FROM users 
WHERE role = $1 
ORDER BY created_at DESC;

-- name: UpdateUser :one
UPDATE users 
SET 
  email = COALESCE($1, email),
  password_hash = COALESCE($2, password_hash),
  role = COALESCE($3, role)
WHERE id = $4
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;