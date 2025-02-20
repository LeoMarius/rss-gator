-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1;



-- name: MarkFeedFetched :exec
UPDATE feeds
SET created_at = CURRENT_TIMESTAMP,
    last_fetched_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: GetNextFeedToFech :one
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC
NULLS FIRST
LIMIT 1;