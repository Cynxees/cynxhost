package tblminecraftserverproperties

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"
	"database/sql"
)

type TblMinecraftServerPropertiesImpl struct {
	DB *sql.DB
}

func New(DB *sql.DB) database.TblMinecraftServerProperties {
	return &TblMinecraftServerPropertiesImpl{
		DB: DB,
	}
}

func (database *TblMinecraftServerPropertiesImpl) InitializeMinecraftServerProperties(ctx context.Context, hostTemplateId int) (context.Context, error) {

	// TODO: implement this
	// SQL := `
	// 	INSERT INTO tbl_minecraft_server_properties (host_template_id, name, value)
	// 	VALUES (?, 'property1', 'value1'), (?, 'property2', 'value2'), (?, 'property3', 'value3')
	// `

	// _, err := database.DB.ExecContext(ctx, SQL, hostTemplateId, hostTemplateId, hostTemplateId)
	// if err != nil {
	// 	return ctx, err
	// }

	return ctx, nil
}

func (database *TblMinecraftServerPropertiesImpl) CreateMinecraftServerProperties(ctx context.Context, data entity.TblMinecraftServerProperties) (context.Context, entity.TblMinecraftServerProperties, error) {
	SQL := `
		INSERT INTO tbl_minecraft_server_properties (host_template_id, name, value)
		VALUES (?, ?, ?)
	`

	result, err := database.DB.ExecContext(ctx, SQL, data.HostTemplateId, data.Name, data.Value)
	if err != nil {
		return ctx, entity.TblMinecraftServerProperties{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return ctx, entity.TblMinecraftServerProperties{}, err
	}

	data.Id = int(id)
	return ctx, data, nil
}

func (database *TblMinecraftServerPropertiesImpl) GetMinecraftServerProperties(ctx context.Context, key, value string) (context.Context, entity.TblMinecraftServerProperties, error) {
	SQL := `
		SELECT id, host_template_id, name, value
		FROM tbl_minecraft_server_properties
		WHERE ` + key + ` = ?
	`

	row := database.DB.QueryRowContext(ctx, SQL, value)

	var properties entity.TblMinecraftServerProperties
	err := row.Scan(&properties.Id, &properties.HostTemplateId, &properties.Name, &properties.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx, entity.TblMinecraftServerProperties{}, nil
		}
		return ctx, entity.TblMinecraftServerProperties{}, err
	}

	return ctx, properties, nil
}

func (database *TblMinecraftServerPropertiesImpl) GetHostTemplateMinecraftServerProperties(ctx context.Context, hostTemplateId int) (context.Context, []entity.TblMinecraftServerProperties, error) {
	SQL := `
		SELECT id, host_template_id, name, value
		FROM tbl_minecraft_server_properties
		WHERE host_template_id = ?
	`

	rows, err := database.DB.QueryContext(ctx, SQL, hostTemplateId)
	if err != nil {
		return ctx, nil, err
	}
	defer rows.Close()

	var propertiesList []entity.TblMinecraftServerProperties
	for rows.Next() {
		var properties entity.TblMinecraftServerProperties
		if err := rows.Scan(&properties.Id, &properties.HostTemplateId, &properties.Name, &properties.Value); err != nil {
			return ctx, nil, err
		}
		propertiesList = append(propertiesList, properties)
	}

	if err := rows.Err(); err != nil {
		return ctx, nil, err
	}

	return ctx, propertiesList, nil
}

func (database *TblMinecraftServerPropertiesImpl) UpdateMinecraftServerProperties(ctx context.Context, data entity.TblMinecraftServerProperties) (context.Context, entity.TblMinecraftServerProperties, error) {
	SQL := `
		UPDATE tbl_minecraft_server_properties
		SET host_template_id = ?, name = ?, value = ?
		WHERE id = ?
	`

	_, err := database.DB.ExecContext(ctx, SQL, data.HostTemplateId, data.Name, data.Value, data.Id)
	if err != nil {
		return ctx, entity.TblMinecraftServerProperties{}, err
	}

	return ctx, data, nil
}

func (database *TblMinecraftServerPropertiesImpl) DeleteMinecraftServerProperties(ctx context.Context, id int) (context.Context, error) {
	SQL := `
		DELETE FROM tbl_minecraft_server_properties
		WHERE id = ?
	`

	_, err := database.DB.ExecContext(ctx, SQL, id)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func (database *TblMinecraftServerPropertiesImpl) DeleteAllByHostTemplateId(ctx context.Context, hostTemplateId int) (context.Context, error) {
	SQL := `
		DELETE FROM tbl_minecraft_server_properties
		WHERE host_template_id = ?
	`

	_, err := database.DB.ExecContext(ctx, SQL, hostTemplateId)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}