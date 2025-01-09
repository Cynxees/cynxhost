-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_instance_type (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    vcpu_count INT NOT NULL,
    memory_size_mb INT NOT NULL,
    spot_price DECIMAL(10, 2) NOT NULL,
    sell_price DECIMAL(10, 2) NOT NULL,
    status ENUM('ACTIVE', 'INACTIVE', 'HIDDEN') NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE tbl_instance (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    aws_instance_id VARCHAR(255) NOT NULL,
    public_ip VARCHAR(255) NOT NULL,
    private_ip VARCHAR(255) NOT NULL,
    instance_type_id INT NOT NULL,
    status VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE tbl_storage (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    size_mb INT NOT NULL,
    aws_ebs_id VARCHAR(255),
    aws_ebs_snapshot_id VARCHAR(255),
    status VARCHAR(255) NOT NULL
);
-- +goose StatementEnd


-- +goose StatementBegin
CREATE TABLE tbl_persistent_node (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    owner_id INT NOT NULL,
    server_template_id INT NOT NULL,
    instance_type_id INT NOT NULL,
    storage_id INT NOT NULL,
    status VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE tbl_instance
ADD CONSTRAINT FK_INSTANCE_INSTANCETYPE
FOREIGN KEY (instance_type_id) REFERENCES tbl_instance_type(id)
ON DELETE RESTRICT
ON UPDATE RESTRICT;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE tbl_persistent_node
ADD CONSTRAINT FK_PERSISTENTNODE_USER
FOREIGN KEY (owner_id) REFERENCES tbl_user(id)
ON DELETE RESTRICT
ON UPDATE RESTRICT;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE tbl_persistent_node
ADD CONSTRAINT FK_PERSISTENTNODE_SERVERTEMPLATE
FOREIGN KEY (server_template_id) REFERENCES tbl_server_template(id)
ON DELETE RESTRICT
ON UPDATE RESTRICT;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE tbl_persistent_node
ADD CONSTRAINT FK_PERSISTENTNODE_INSTANCETYPE
FOREIGN KEY (instance_type_id) REFERENCES tbl_instance_type(id)
ON DELETE RESTRICT
ON UPDATE RESTRICT;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE tbl_persistent_node
ADD CONSTRAINT FK_PERSISTENTNODE_STORAGE
FOREIGN KEY (storage_id) REFERENCES tbl_storage(id)
ON DELETE RESTRICT
ON UPDATE RESTRICT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tbl_persistent_node DROP FOREIGN KEY FK_PERSISTENTNODE_USER;
ALTER TABLE tbl_persistent_node DROP FOREIGN KEY FK_PERSISTENTNODE_SERVERTEMPLATE;
ALTER TABLE tbl_persistent_node DROP FOREIGN KEY FK_PERSISTENTNODE_INSTANCETYPE;
ALTER TABLE tbl_persistent_node DROP FOREIGN KEY FK_PERSISTENTNODE_STORAGE;
ALTER TABLE tbl_instance DROP FOREIGN KEY FK_INSTANCE_INSTANCETYPE;

DROP TABLE IF EXISTS tbl_persistent_node;
DROP TABLE IF EXISTS tbl_storage;
DROP TABLE IF EXISTS tbl_instance;
DROP TABLE IF EXISTS tbl_instance_type;
-- +goose StatementEnd
