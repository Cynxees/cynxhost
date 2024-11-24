-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_ami (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  image_id VARCHAR(255) NOT NULL,
  minecraft_edition VARCHAR(255) NOT NULL,
  minecraft_version VARCHAR(255) NOT NULL,
  mod_loader VARCHAR(255) NOT NULL,
  mod_loader_version VARCHAR(255) NOT NULL,
  minimum_ram INT NOT NULL,
  minimum_vcpu INT NOT NULL,
  created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  modified_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tbl_ami
-- +goose StatementEnd
