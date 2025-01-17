-- name: CreateFeedFollow :one
WITH inserted_feed_follows AS (
    INSERT INTO feed_follows (id, user_id, feed_id, created_at, updated_at)
    VALUES (
        $1, $2, $3, $4, $5
    )
    RETURNING *
)
SELECT
    inserted_feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follows
JOIN feeds ON inserted_feed_follows.feed_id = feeds.id
JOIN users ON inserted_feed_follows.user_id = users.id;

-- name: GetFeedFollowsForUser :many
SELECT
    ff.id,
    ff.user_id,
    ff.feed_id,
    ff.created_at,
    ff.updated_at,
    f.name AS feed_name,
    u.name AS user_name
FROM feed_follows AS ff
JOIN users AS u ON ff.user_id = u.id
JOIN feeds AS f ON ff.feed_id = f.id
WHERE ff.user_id = $1;

-- name: DeleteFeedByUrlForUser :exec
WITH feed_id AS (
    SELECT feeds.id FROM feeds WHERE feeds.url = $1
)
DELETE FROM feed_follows
WHERE feed_follows.feed_id = feed_id
AND feed_follows.user_id = $2;