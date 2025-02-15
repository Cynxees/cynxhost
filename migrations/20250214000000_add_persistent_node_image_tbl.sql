-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_persistent_node_image (
  id INT AUTO_INCREMENT PRIMARY KEY,
  persistent_node_id int NOT NULL UNIQUE,
  image_tag VARCHAR(255),
  status VARCHAR(255) NOT NULL,
  created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE tbl_persistent_node_image
ADD CONSTRAINT FK_PERSISTENTNODEIMAGE_PERSISTENTNODE
FOREIGN KEY (persistent_node_id) REFERENCES tbl_persistent_node(id)
ON DELETE RESTRICT
ON UPDATE RESTRICT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tbl_persistent_node_image DROP FOREIGN KEY FK_PERSISTENTNODEIMAGE_PERSISTENTNODE;
DROP TABLE IF EXISTS tbl_persistent_node_image;
-- +goose StatementEnd
