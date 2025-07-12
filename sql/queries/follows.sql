-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT
    iff.id,
    iff.created_at,
    iff.updated_at,
    iff.user_id,
    iff.feed_id,
    u.name AS user_name,
    f.name AS feed_name
FROM inserted_feed_follow AS iff
JOIN users AS u ON iff.user_id = u.id
JOIN feeds AS f ON iff.feed_id = f.id;

-- name: ListFeedFollow :many
SELECT
    ff.id,
    ff.created_at,
    ff.updated_at,
    ff.user_id,
    ff.feed_id,
    u.name AS user_name,
    f.name AS feed_name
FROM feed_follows AS ff
JOIN users AS u ON ff.user_id = u.id
JOIN feeds AS f ON ff.feed_id = f.id
WHERE ff.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE user_id = $1 and feed_id = $2;
