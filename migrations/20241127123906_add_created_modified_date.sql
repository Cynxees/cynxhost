-- +goose Up
-- +goose StatementBegin
ALTER TABLE tbl_instance ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, ADD COLUMN modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tbl_instance DROP COLUMN created_at, DROP COLUMN modified_at;
-- +goose StatementEnd
