-- name: CreatePickupPoint :one
INSERT INTO pickup_points (
  name,
  city,
  address
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetPickupPoint :one
SELECT * FROM pickup_points 
WHERE id = $1 LIMIT 1;

-- name: ListPickupPointsByCity :many
SELECT * FROM pickup_points 
WHERE city = $1 
ORDER BY registered_at DESC;

-- name: UpdatePickupPoint :one
UPDATE pickup_points 
SET 
  name = COALESCE($1, name),
  address = COALESCE($2, address)
WHERE id = $3
RETURNING *;

-- name: DeletePickupPoint :exec
DELETE FROM pickup_points
WHERE id = $1
RETURNING id;