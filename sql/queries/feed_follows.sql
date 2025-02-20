
-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2)
    RETURNING *
)
SELECT
    inserted_feed_follow.*,
    u.name AS user_name,
    f.name AS feed_name
FROM inserted_feed_follow
INNER JOIN users AS u ON inserted_feed_follow.user_id = u.id
INNER JOIN feeds AS f ON inserted_feed_follow.feed_id = f.id;


-- name: GetFeedFollowsForUser :many
SELECT feeds.* 
FROM feeds
INNER JOIN feed_follows ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1;