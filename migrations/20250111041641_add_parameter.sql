-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_parameter (
    id VARCHAR(255) PRIMARY KEY,
    value TEXT NOT NULL,
    description TEXT,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tbl_parameter;
-- +goose StatementEnd
