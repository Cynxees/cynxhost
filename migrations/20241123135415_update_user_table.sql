-- +goose Up
-- +goose StatementBegin
ALTER TABLE tbl_user
ADD COLUMN created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN modified_date DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tbl_user
DROP COLUMN created_date,
DROP COLUMN modified_date;
-- +goose StatementEnd
