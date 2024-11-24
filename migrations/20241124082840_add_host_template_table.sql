-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_host_template (
  id INT AUTO_INCREMENT PRIMARY KEY,
  
  owner_id INT NOT NULL,
  ami_id INT NOT NULL,
  instance_type_id INT NOT NULL,
  
  created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  modified_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  
  FOREIGN KEY (owner_id) REFERENCES tbl_user(id),
  FOREIGN KEY (ami_id) REFERENCES tbl_ami(id),
  FOREIGN KEY (instance_type_id) REFERENCES tbl_instance_type(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tbl_host_template;
-- +goose StatementEnd
