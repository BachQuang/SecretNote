-- name: CreatePost :one
INSERT INTO posts (
    email,
    title,
    content
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: ListPosts :many
SELECT * FROM posts
WHERE 
    email = $1
ORDER BY created_at desc
LIMIT $2
OFFSET $3;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;