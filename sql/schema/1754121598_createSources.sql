-- +goose Up
CREATE TABLE argumentSources (
  id UUID PRIMARY KEY,
  creation_date TIMESTAMP NOT NULL,
  last_update_time TIMESTAMP NOT NULL,
  content TEXT,
  argument_id UUID NOT NULL,
  FOREIGN KEY (argument_id)
    REFERENCES argument(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE argumentSources;
