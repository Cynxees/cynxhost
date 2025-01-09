-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_script (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    variables JSON NOT NULL,
    setup_script TEXT NOT NULL,
    start_script TEXT NOT NULL,
    stop_script TEXT NOT NULL,
    shutdown_script TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE tbl_server_template (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    minimum_ram INT NOT NULL,
    minimum_cpu INT NOT NULL,
    minimum_disk INT NOT NULL,
    script_id INT NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE tbl_server_template
ADD CONSTRAINT FK_SERVERTEMPLATE_SCRIPT
FOREIGN KEY (script_id) REFERENCES tbl_script(id)
ON DELETE RESTRICT
ON UPDATE RESTRICT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tbl_server_template DROP FOREIGN KEY FK_SERVERTEMPLATE_SCRIPT;
DROP TABLE IF EXISTS tbl_server_template;
DROP TABLE IF EXISTS tbl_script;
-- +goose StatementEnd
