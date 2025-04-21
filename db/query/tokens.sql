-- name: CreateToken :one
INSERT INTO tokens (
  user_id,
  token,
  expires_at
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetToken :one
SELECT * FROM tokens 
WHERE token = $1 LIMIT 1;

-- name: DeleteExpiredTokens :exec
DELETE FROM tokens 
WHERE expires_at < CURRENT_TIMESTAMP 
RETURNING *;

-- name: UpdateTokenExpiry :one
UPDATE tokens 
SET expires_at = $1
WHERE token = $2 AND expires_at > CURRENT_TIMESTAMP
RETURNING *;
