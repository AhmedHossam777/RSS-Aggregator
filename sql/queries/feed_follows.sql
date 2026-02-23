-- name: CreateFeedFollow :one
INSERT INTO feed_follows (ID,user_id,feed_id) VALUES ($1,$2,$3)
returning *;

-- name: GetFeedFollows :many
SELECT * FROM feed_follows WHERE user_id = $1;

-- name: DeleteFeedFollow :one
DELETE FROM feed_follows WHERE feed_id = $1 AND user_id = $2
returning *;