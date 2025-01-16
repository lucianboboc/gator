-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id, created_at, updated_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT f.id, f.name, f.url, f.user_id, f.created_at, u.name AS user_name, u.created_at, u.updated_at FROM feeds AS f
JOIN users AS u ON u.id = f.user_id;