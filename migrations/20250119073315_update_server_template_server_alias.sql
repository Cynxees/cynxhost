-- +goose Up
-- +goose StatementBegin
ALTER TABLE tbl_server_template ADD COLUMN image_path VARCHAR(255);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE tbl_persistent_node ADD COLUMN server_alias VARCHAR(255) NOT NULL UNIQUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tbl_server_template DROP COLUMN image_path;
ALTER TABLE tbl_persistent_node DROP COLUMN server_alias;
-- +goose StatementEnd
