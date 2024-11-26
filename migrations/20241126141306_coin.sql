-- +goose Up
-- +goose StatementBegin
ALTER TABLE tbl_user 
    ADD COLUMN coin INT NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE tbl_ami 
    MODIFY mod_loader VARCHAR(255);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE tbl_ami 
    MODIFY mod_loader_version VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
