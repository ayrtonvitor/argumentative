-- +goose Up
CREATE TABLE argument (
    id UUID PRIMARY KEY,
    creation_date TIMESTAMP NOT NULL,
    last_update_time TIMESTAMP NOT NULL,
    brief TEXT NOT NULL,
    description TEXT
);

-- +goose Down
DROP TABLE argument;
