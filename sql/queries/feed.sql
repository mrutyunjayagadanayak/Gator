-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeed :many
SELECT * FROM feeds
WHERE user_id = $1;

-- name: GetAllFeeds :many 
SELECT * FROM feeds;

-- name: GetFeedByURL :one 
SELECT * FROM feeds
WHERE url = $1;

-- name: UpdateFeedFetchTime :exec
UPDATE feeds
SET last_fetched_at = $1
WHERE id = $2;

-- name: GetNextFeedToFetch :one 
SELECT * FROM feeds
WHERE last_fetched_at IS NULL OR last_fetched_at < NOW() - INTERVAL '10 seconds'
ORDER BY last_fetched_at NULLS FIRST
LIMIT 1;
