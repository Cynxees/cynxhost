-- +goose Up
-- +goose StatementBegin
ALTER TABLE tbl_server_template_category 
ADD COLUMN image_path VARCHAR(255);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE tbl_instance_type
ADD COLUMN image_path VARCHAR(255),
ADD COLUMN description TEXT NOT NULL,
ADD COLUMN aws_key VARCHAR(255) NOT NULL UNIQUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tbl_server_template_category 
DROP COLUMN image_path;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE tbl_instance_type
DROP COLUMN image_path,
DROP COLUMN description,
DROP COLUMN aws_key;
-- +goose StatementEnd
