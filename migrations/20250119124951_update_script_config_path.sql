-- +goose Up
-- +goose StatementBegin
ALTER TABLE tbl_script
ADD COLUMN config_path VARCHAR(255) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tbl_script
DROP COLUMN config_path;
-- +goose StatementEnd
