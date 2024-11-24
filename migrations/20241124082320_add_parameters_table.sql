-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_parameters (
  id INT AUTO_INCREMENT PRIMARY KEY,
  key VARCHAR(255) NOT NULL,
  value TEXT NOT NULL,
  description TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tbl_parameters
-- +goose StatementEnd
