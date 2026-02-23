-- name: CreateFeedFollow :one
INSERT INTO feed_follows (ID,user_id,feed_id) VALUES ($1,$2,$3)
returning *;

-- name: GetFeedFollows :many
SELECT * FROM feed_follows WHERE user_id = $1;