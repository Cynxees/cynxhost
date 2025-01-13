-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_server_template_category (
    id int AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    parent_id INT,
    server_template_id INT,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE tbl_server_template_category
ADD CONSTRAINT FK_SERVERTEMPLATECATEGORY_SERVERTEMPLATE
FOREIGN KEY (server_template_id) REFERENCES tbl_server_template(id)
ON DELETE RESTRICT
ON UPDATE RESTRICT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tbl_server_template_category DROP FOREIGN KEY FK_SERVERTEMPLATECATEGORY_SERVERTEMPLATE;
DROP TABLE IF EXISTS tbl_server_template_category;
-- +goose StatementEnd
