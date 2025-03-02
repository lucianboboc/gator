// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
WITH inserted_feed_follows AS (
    INSERT INTO feed_follows (id, user_id, feed_id, created_at, updated_at)
    VALUES (
        $1, $2, $3, $4, $5
    )
    RETURNING id, user_id, feed_id, created_at, updated_at
)
SELECT
    inserted_feed_follows.id, inserted_feed_follows.user_id, inserted_feed_follows.feed_id, inserted_feed_follows.created_at, inserted_feed_follows.updated_at,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follows
JOIN feeds ON inserted_feed_follows.feed_id = feeds.id
JOIN users ON inserted_feed_follows.user_id = users.id
`

type CreateFeedFollowParams struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	FeedID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateFeedFollowRow struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	FeedID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedName  string
	UserName  string
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (CreateFeedFollowRow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.UserID,
		arg.FeedID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i CreateFeedFollowRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FeedID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FeedName,
		&i.UserName,
	)
	return i, err
}

const deleteFeedByUrlForUser = `-- name: DeleteFeedByUrlForUser :exec
WITH feed_id AS (
    SELECT feeds.id FROM feeds WHERE feeds.url = $1
)
DELETE FROM feed_follows
WHERE feed_follows.feed_id = feed_id
AND feed_follows.user_id = $2
`

type DeleteFeedByUrlForUserParams struct {
	Url    string
	UserID uuid.UUID
}

func (q *Queries) DeleteFeedByUrlForUser(ctx context.Context, arg DeleteFeedByUrlForUserParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedByUrlForUser, arg.Url, arg.UserID)
	return err
}

const getFeedFollowsForUser = `-- name: GetFeedFollowsForUser :many
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
WHERE ff.user_id = $1
`

type GetFeedFollowsForUserRow struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	FeedID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedName  string
	UserName  string
}

func (q *Queries) GetFeedFollowsForUser(ctx context.Context, userID uuid.UUID) ([]GetFeedFollowsForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedFollowsForUserRow
	for rows.Next() {
		var i GetFeedFollowsForUserRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.FeedID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.FeedName,
			&i.UserName,
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
