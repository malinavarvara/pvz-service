// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  email,
  password_hash,
  role
) VALUES (
  $1, $2, $3
) RETURNING id, email, password_hash, role, created_at
`

type CreateUserParams struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.PasswordHash, arg.Role)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password_hash, role, created_at FROM users 
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const listUsersByRole = `-- name: ListUsersByRole :many
SELECT id, email, password_hash, role, created_at FROM users 
WHERE role = $1 
ORDER BY created_at DESC
`

func (q *Queries) ListUsersByRole(ctx context.Context, role string) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsersByRole, role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.PasswordHash,
			&i.Role,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users 
SET 
  email = COALESCE($1, email),
  password_hash = COALESCE($2, password_hash),
  role = COALESCE($3, role)
WHERE id = $4
RETURNING id, email, password_hash, role, created_at
`

type UpdateUserParams struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
	ID           int32  `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Email,
		arg.PasswordHash,
		arg.Role,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}
