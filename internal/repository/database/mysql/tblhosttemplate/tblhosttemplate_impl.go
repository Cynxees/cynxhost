package tblhosttemplate

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"
	"database/sql"
	"strconv"
)

type TblHostTemplateImpl struct {
	DB *sql.DB
}

func New(DB *sql.DB) database.TblHostTemplate {
	return &TblHostTemplateImpl{
		DB: DB,
	}
}

func (database *TblHostTemplateImpl) CreateHostTemplate(ctx context.Context, hostTemplate entity.TblHostTemplate) (context.Context, entity.TblHostTemplate, error) {
	SQL := `
		INSERT INTO tbl_host_template (owner_id, ami_id, instance_type_id)
		VALUES (?, ?, ?)
	`

	result, err := database.DB.ExecContext(ctx, SQL, hostTemplate.OwnerId, hostTemplate.AmiId, hostTemplate.InstanceTypeId)
	if err != nil {
		return ctx, entity.TblHostTemplate{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return ctx, entity.TblHostTemplate{}, err
	}

	ctx, createdData, err := database.GetHostTemplate(ctx, "id", strconv.Itoa(int(id)))
	if err != nil {
		return ctx, entity.TblHostTemplate{}, err
	}
	
	return ctx, createdData, nil
}

func (database *TblHostTemplateImpl) GetHostTemplate(ctx context.Context, key, value string) (context.Context, entity.TblHostTemplate, error) {
	SQL := `
		SELECT id, owner_id, ami_id, instance_type_id
		FROM tbl_host_template
		WHERE ` + key + ` = ?
	`

	row := database.DB.QueryRowContext(ctx, SQL, value)

	var hostTemplate entity.TblHostTemplate
	err := row.Scan(&hostTemplate.Id, &hostTemplate.OwnerId, &hostTemplate.AmiId, &hostTemplate.InstanceTypeId)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx, entity.TblHostTemplate{}, nil
		}
		return ctx, entity.TblHostTemplate{}, err
	}

	return ctx, hostTemplate, nil
}

func (database *TblHostTemplateImpl) GetAllUserOwnedHostTemplate(ctx context.Context, userId int) (context.Context, []entity.TblHostTemplate, error) {
	SQL := `
		SELECT id, owner_id, ami_id, instance_type_id
		FROM tbl_host_template
		WHERE owner_id = ?
	`

	rows, err := database.DB.QueryContext(ctx, SQL, userId)
	if err != nil {
		return ctx, nil, err
	}
	defer rows.Close()

	var hostTemplates []entity.TblHostTemplate
	for rows.Next() {
		var hostTemplate entity.TblHostTemplate
		if err := rows.Scan(&hostTemplate.Id, &hostTemplate.OwnerId, &hostTemplate.AmiId, &hostTemplate.InstanceTypeId); err != nil {
			return ctx, nil, err
		}
		hostTemplates = append(hostTemplates, hostTemplate)
	}

	if err := rows.Err(); err != nil {
		return ctx, nil, err
	}

	return ctx, hostTemplates, nil
}

func (database *TblHostTemplateImpl) UpdateHostTemplate(ctx context.Context, hostTemplate entity.TblHostTemplate) (context.Context, entity.TblHostTemplate, error) {
	SQL := `
		UPDATE tbl_host_template
		SET owner_id = ?, ami_id = ?, instance_type_id = ?
		WHERE id = ?
	`

	result, err := database.DB.ExecContext(ctx, SQL, hostTemplate.OwnerId, hostTemplate.AmiId, hostTemplate.InstanceTypeId, hostTemplate.Id)
	if err != nil {
		return ctx, entity.TblHostTemplate{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return ctx, entity.TblHostTemplate{}, err
	}

	ctx, finalData, err := database.GetHostTemplate(ctx, "id", strconv.Itoa(hostTemplate.Id))
	if err != nil {
		return ctx, entity.TblHostTemplate{}, err
	}

	return ctx, finalData, nil
}

func (database *TblHostTemplateImpl) DeleteHostTemplate(ctx context.Context, hostTemplateId int) (context.Context, error) {
	SQL := `
		DELETE FROM tbl_host_template
		WHERE id = ?
	`

	_, err := database.DB.ExecContext(ctx, SQL, hostTemplateId)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}
