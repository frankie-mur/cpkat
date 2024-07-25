-- name: GetLink :one
SELECT * FROM links 
WHERE id = ? LIMIT 1;

-- name: ListLinks :many
SELECT * FROM links
ORDER BY name;

-- name: CreateLink :one
INSERT INTO links(
  name, link
) VALUES (
  ?, ?
)
RETURNING *;

-- name: UpdateLink :exec
UPDATE links
set name = ?,
link = ?
WHERE id = ?;

-- name: DeleteAuthor :exec
DELETE FROM links
WHERE id = ?;
