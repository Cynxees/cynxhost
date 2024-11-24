-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_instance_type (
  id INT AUTO_INCREMENT PRIMARY KEY,
  vcpu_count INT NOT NULL,
  memory_count INT NOT NULL,
  spot_price INT NOT NULL,
  sell_price INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  status VARCHAR(50) NOT NULL,
  created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  modified_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tbl_instance_type
-- +goose StatementEnd
