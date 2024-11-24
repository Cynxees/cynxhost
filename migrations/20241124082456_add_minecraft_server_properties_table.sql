-- +goose Up
CREATE TABLE tbl_minecraft_server_properties (
  Key VARCHAR(255) PRIMARY KEY NOT NULL,
  Value VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tbl_minecraft_server_properties
-- +goose StatementEnd
