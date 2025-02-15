-- +goose Up
-- +goose StatementBegin

CREATE TABLE tbl_persistent_node_image (
  id INT AUTO_INCREMENT PRIMARY KEY,
  persistent_node_id VARCHAR(255 NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  coin INT NOT NULL DEFAULT 0,
  created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tbl_server_template
DROP COLUMN image_url,
ADD COLUMN image_path VARCHAR(255) NOT NULL;
-- +goose StatementEnd
