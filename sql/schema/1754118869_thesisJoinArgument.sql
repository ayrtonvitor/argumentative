-- +goose Up
CREATE TABLE thesis_argument (
  thesis_id UUID NOT NULL,
  argument_id UUID NOT NULL,
  PRIMARY KEY (thesis_id, argument_id),
  FOREIGN KEY (thesis_id)
    REFERENCES thesis(id)
    ON DELETE CASCADE,
  FOREIGN KEY (argument_id)
    REFERENCES argument(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE thesis_argument;
