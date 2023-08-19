-- name: CreateFeedFollow :one
INSERT INTO feed_follow (id ,feed_id ,user_id ,created_at ,updated_at )
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follow WHERE id = $1;

-- name: GetFeedFollowForUser :many
SELECT * FROM feed_follow WHERE user_id = $1;