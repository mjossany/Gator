-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeedsWithUsers :many
SELECT
    feeds.id,
    feeds.name,
    feeds.url,
    feeds.created_at,
    feeds.updated_at,
    users.name AS user_name
FROM feeds
JOIN users ON feeds.user_id = users.id
ORDER BY feeds.created_at DESC;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;