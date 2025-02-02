-- +goose Up
-- +goose StatementBegin
ALTER TABLE tbl_server_template
DROP COLUMN image_path,
ADD COLUMN image_url VARCHAR(255) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tbl_server_template
DROP COLUMN image_url,
ADD COLUMN image_path VARCHAR(255) NOT NULL;
-- +goose StatementEnd
