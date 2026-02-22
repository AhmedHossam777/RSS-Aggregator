-- +goose up
CREATE TABLE feeds(
    id UUID PRIMARY KEY ,
    name TEXT NOT NULL ,
    created_at TIMESTAMP NOT NULL ,
    updated_at TIMESTAMP NOT NULL ,
    url TEXT UNIQUE NOT NULL,
    user_id UUID REFERENCES users(id) on DELETE CASCADE
);

-- +goose down
DROP TABLE feeds;