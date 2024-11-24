package tblinstance

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"
	"database/sql"
	"strconv"
)

type TblInstanceImpl struct {
	DB *sql.DB
}

func New(DB *sql.DB) database.TblInstance {
	return &TblInstanceImpl{
		DB: DB,
	}
}
func (database *TblInstanceImpl) CreateInstance(ctx context.Context, instance entity.TblInstance) (context.Context, entity.TblInstance, error) {
	SQL := `
		INSERT INTO tbl_instance (host_template_id, owner_id, status, address)
		VALUES (?, ?, ?, ?)
	`

	result, err := database.DB.ExecContext(ctx, SQL, instance.HostTemplateId, instance.OwnerId, instance.Status, instance.Address)
	if err != nil {
		return ctx, entity.TblInstance{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return ctx, entity.TblInstance{}, err
	}

	ctx, createdData, err := database.GetInstance(ctx, "id", strconv.Itoa(int(id)))
	if err != nil {
		return ctx, entity.TblInstance{}, err
	}

	return ctx, createdData, nil
}

func (database *TblInstanceImpl) GetInstance(ctx context.Context, key, value string) (context.Context, entity.TblInstance, error) {
	SQL := `
		SELECT id, host_template_id, owner_id, status, address
		FROM tbl_instance
		WHERE ` + key + ` = ?
	`

	row := database.DB.QueryRowContext(ctx, SQL, value)

	var instance entity.TblInstance
	err := row.Scan(&instance.Id, &instance.HostTemplateId, &instance.OwnerId, &instance.Status, &instance.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx, entity.TblInstance{}, nil
		}
		return ctx, entity.TblInstance{}, err
	}

	return ctx, instance, nil
}

func (database *TblInstanceImpl) UpdateInstance(ctx context.Context, instance entity.TblInstance) (context.Context, entity.TblInstance, error) {
	SQL := `
		UPDATE tbl_instance
		SET host_template_id = ?, owner_id = ?, status = ?, address = ?
		WHERE id = ?
	`

	result, err := database.DB.ExecContext(ctx, SQL, instance.HostTemplateId, instance.OwnerId, instance.Status, instance.Address, instance.Id)
	if err != nil {
		return ctx, entity.TblInstance{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return ctx, entity.TblInstance{}, err
	}

	ctx, updatedData, err := database.GetInstance(ctx, "id", strconv.Itoa(instance.Id))
	if err != nil {
		return ctx, entity.TblInstance{}, err
	}

	return ctx, updatedData, nil
}

func (database *TblInstanceImpl) DeleteInstance(ctx context.Context, instanceId int) (context.Context, error) {
	SQL := `
		DELETE FROM tbl_instance
		WHERE id = ?
	`

	_, err := database.DB.ExecContext(ctx, SQL, instanceId)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}