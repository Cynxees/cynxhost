-- +goose Up
-- +goose StatementBegin
ALTER TABLE tbl_instance_type
DROP COLUMN memory_size_mb,
ADD COLUMN memory_size_gb INTEGER,
ADD COLUMN network_speed_mbps INTEGER;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tbl_instance_type
DROP COLUMN memory_size_gb,
DROP COLUMN network_speed_mbps,
ADD COLUMN memory_size_mb INTEGER;
-- +goose StatementEnd
