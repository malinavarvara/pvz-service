-- name: AddProduct :one
INSERT INTO products (
  receiving_id,
  type,
  description,
  avito_order_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetLastProduct :one
SELECT * FROM products 
WHERE receiving_id = $1 
ORDER BY added_at DESC 
LIMIT 1;

-- name: ListProductsInReceiving :many
SELECT * FROM products 
WHERE receiving_id = $1 
ORDER BY added_at ASC;

-- name: UpdateProduct :one
UPDATE products AS p
SET
    type = COALESCE($1, p.type),
    description = COALESCE($2, p.description)
WHERE p.id = $3 AND p.receiving_id IN (
    SELECT r.id FROM receivings AS r
    WHERE r.status = 'in_progress' AND r.pickup_point_id = $4
)
RETURNING *;

-- name: DeleteLastProduct :one
DELETE FROM products AS p
WHERE p.id = (
  SELECT id FROM products AS p2
  WHERE p2.receiving_id = $1
  ORDER BY p2.added_at DESC
  LIMIT 1
)
RETURNING *;