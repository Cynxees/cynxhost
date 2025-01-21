-- +goose Up
-- +goose StatementBegin
ALTER TABLE tbl_persistent_node
ADD COLUMN variables JSON NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tbl_persistent_node
DROP COLUMN variables;
-- +goose StatementEnd
