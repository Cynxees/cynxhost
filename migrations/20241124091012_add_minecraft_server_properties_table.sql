-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_minecraft_server_properties (
  id INT AUTO_INCREMENT PRIMARY KEY,
  host_template_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  value VARCHAR(255) NOT NULL,

  FOREIGN KEY (host_template_id) REFERENCES tbl_host_template(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tbl_minecraft_server_properties;
-- +goose StatementEnd