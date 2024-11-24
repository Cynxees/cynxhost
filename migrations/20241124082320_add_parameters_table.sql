-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_parameters (
  name VARCHAR(255) NOT NULL PRIMARY KEY,
  value TEXT NOT NULL,
  description TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tbl_parameters
-- +goose StatementEnd
