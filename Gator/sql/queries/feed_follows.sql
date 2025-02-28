-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows  (created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4
    )
    RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users ON users.id = inserted_feed_follow.user_id
INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT
    feeds.name AS feed_name
FROM feed_follows
INNER JOIN feeds ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollows :exec
DELETE 
FROM feed_follows
WHERE user_id = $1 AND feed_id = $2;