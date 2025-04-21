-- name: StartReceiving :one
INSERT INTO receivings (
  pickup_point_id,
  status
) VALUES (
  $1, 'in_progress'
) RETURNING *;

-- name: CreateReceiving :one
INSERT INTO receivings (pickup_point_id, status)
VALUES ($1, $2)
RETURNING *;

-- name: CloseReceiving :one
UPDATE receivings 
SET 
  status = 'closed',
  closed_at = CURRENT_TIMESTAMP
WHERE id = $1 
RETURNING *;

-- name: GetActiveReceiving :one
SELECT * FROM receivings 
WHERE pickup_point_id = $1 AND status = 'in_progress'
LIMIT 1;

-- name: UpdateReceivingStatus :one
UPDATE receivings 
SET
  status = $1::varchar, -- Явное указание типа
  closed_at = CASE WHEN $1::varchar = 'closed' THEN NOW() ELSE NULL END
WHERE id = $2::bigint  -- Явное указание типа для ID
RETURNING *;

-- name: DeleteEmptyReceiving :exec
DELETE FROM receivings AS r
WHERE r.id = $1 AND r.status = 'in_progress'
AND NOT EXISTS (
  SELECT 1 FROM products WHERE receiving_id = $1
);

-- name: DeleteAllProductsFromReceiving :exec
DELETE FROM products
WHERE receiving_id = $1;