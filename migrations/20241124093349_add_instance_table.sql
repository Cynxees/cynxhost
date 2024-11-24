-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_instance (
  id INT AUTO_INCREMENT PRIMARY KEY,
  host_template_id INT NOT NULL,
  owner_id INT NOT NULL,
  status VARCHAR(255) NOT NULL,
  address VARCHAR(255) NOT NULL,

  FOREIGN KEY (host_template_id) REFERENCES tbl_host_template(id),
  FOREIGN KEY (owner_id) REFERENCES tbl_user(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tbl_instance
-- +goose StatementEnd
