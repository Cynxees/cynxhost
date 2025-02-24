-- +goose Up
-- +goose StatementBegin
ALTER TABLE tbl_persistent_node
ADD COLUMN contain_public_ipv4 BOOLEAN NOT NULL DEFAULT FALSE;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE tbl_instance
ADD COLUMN public_ipv6 VARCHAR(255),
ADD COLUMN private_ipv6 VARCHAR(255),
MODIFY public_ip VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tbl_persistent_node
DROP COLUMN memory_size_gb,
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE tbl_instance
DROP COLUMN public_ipv6,
DROP COLUMN private_ipv6;
-- +goose StatementEnd
