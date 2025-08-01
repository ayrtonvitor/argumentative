-- +goose Up
CREATE TABLE thesis (
    id UUID PRIMARY KEY,
    creation_date TIMESTAMP NOT NULL,
    last_update_time TIMESTAMP NOT NULL,
    title TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE thesis;
