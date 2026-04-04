-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name as feed_name,
    users.name as user_name
FROM inserted_feed_follow
INNER JOIN users on user_id = users.id
INNER JOIN feeds on feed_id = feeds.id;

-- name: ListFeedFollowsForUser :many
SELECT feed_follows.*, feeds.name as feed_name, users.name as user_name
FROM feed_follows
INNER JOIN users on user_id = users.id
INNER JOIN feeds on feed_id = feeds.id
WHERE feed_follows.user_id = $1;
