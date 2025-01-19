-- +goose Up
-- +goose StatementBegin
ALTER TABLE tbl_persistent_node
ADD COLUMN dns_record_id VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tbl_persistent_node DROP COLUMN dns_record_id;
-- +goose StatementEnd
