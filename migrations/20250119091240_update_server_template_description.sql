-- +goose Up
-- +goose StatementBegin
ALTER TABLE tbl_server_template
ADD COLUMN description TEXT;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE tbl_server_template_category
ADD COLUMN description TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tbl_server_template_category
DROP COLUMN description;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE tbl_server_template
DROP COLUMN description;
-- +goose StatementEnd
