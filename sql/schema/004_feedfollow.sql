-- +goose Up
CREATE TABLE feed_follow (
    id UUID PRIMARY KEY,
    feed_id UUID NOT NULL REFERENCES feeds(id),
    user_id UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
    
);

-- +goose Down
DROP TABLE feed_follow;