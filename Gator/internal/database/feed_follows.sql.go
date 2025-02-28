// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows  (created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4
    )
    RETURNING id, created_at, updated_at, user_id, feed_id
)
SELECT
    inserted_feed_follow.id, inserted_feed_follow.created_at, inserted_feed_follow.updated_at, inserted_feed_follow.user_id, inserted_feed_follow.feed_id,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users ON users.id = inserted_feed_follow.user_id
INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id
`

type CreateFeedFollowParams struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int32
	FeedID    int32
}

type CreateFeedFollowRow struct {
	ID        int32
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int32
	FeedID    int32
	FeedName  string
	UserName  string
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (CreateFeedFollowRow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i CreateFeedFollowRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
		&i.FeedName,
		&i.UserName,
	)
	return i, err
}

const deleteFeedFollows = `-- name: DeleteFeedFollows :exec
DELETE 
FROM feed_follows
WHERE user_id = $1 AND feed_id = $2
`

type DeleteFeedFollowsParams struct {
	UserID int32
	FeedID int32
}

func (q *Queries) DeleteFeedFollows(ctx context.Context, arg DeleteFeedFollowsParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollows, arg.UserID, arg.FeedID)
	return err
}

const getFeedFollowsForUser = `-- name: GetFeedFollowsForUser :many
SELECT
    feeds.name AS feed_name
FROM feed_follows
INNER JOIN feeds ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
`

func (q *Queries) GetFeedFollowsForUser(ctx context.Context, userID int32) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var feed_name string
		if err := rows.Scan(&feed_name); err != nil {
			return nil, err
		}
		items = append(items, feed_name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
