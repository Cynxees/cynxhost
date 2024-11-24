package tblinstancetype

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"
	"database/sql"
)

type TblInstanceTypeImpl struct {
	DB *sql.DB
}

func New(DB *sql.DB) database.TblInstanceType {
	return &TblInstanceTypeImpl{
		DB: DB,
	}
}

func (database *TblInstanceTypeImpl) GetInstanceType(ctx context.Context, key, value string) (context.Context, entity.TblInstanceType, error) {
	query := "SELECT id, vcpu_count, memory_count, spot_price, sell_price, name, status, created_date, modified_date FROM tbl_instance_type WHERE " + key + " = ?"
	row := database.DB.QueryRowContext(ctx, query, value)

	var instance entity.TblInstanceType
	err := row.Scan(&instance.Id, &instance.VcpuCount, &instance.MemoryCount, &instance.SpotPrice, &instance.SellPrice, &instance.Name, &instance.Status, &instance.CreatedDate, &instance.ModifiedDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx, entity.TblInstanceType{}, nil
		}
		return ctx, entity.TblInstanceType{}, err
	}

	return ctx, instance, nil
}

func (database *TblInstanceTypeImpl) GetAllInstanceType(ctx context.Context) (context.Context, []entity.TblInstanceType, error) {
	query := "SELECT id, vcpu_count, memory_count, spot_price, sell_price, name, status, created_date, modified_date FROM tbl_instance_type"
	rows, err := database.DB.QueryContext(ctx, query)
	if err != nil {
		return ctx, nil, err
	}
	defer rows.Close()

	var instances []entity.TblInstanceType
	for rows.Next() {
		var instance entity.TblInstanceType
		err := rows.Scan(&instance.Id, &instance.VcpuCount, &instance.MemoryCount, &instance.SpotPrice, &instance.SellPrice, &instance.Name, &instance.Status, &instance.CreatedDate, &instance.ModifiedDate)
		if err != nil {
			return ctx, nil, err
		}
		instances = append(instances, instance)
	}

	if err = rows.Err(); err != nil {
		return ctx, nil, err
	}

	return ctx, instances, nil
}

func (database *TblInstanceTypeImpl) UpdateInstanceType(ctx context.Context, instance entity.TblInstanceType) (context.Context, entity.TblInstanceType, error) {
	query := "UPDATE tbl_instance_type SET vcpu_count = ?, memory_count = ?, spot_price = ?, sell_price = ?, name = ?, status = ? WHERE id = ?"
	_, err := database.DB.ExecContext(ctx, query, instance.VcpuCount, instance.MemoryCount, instance.SpotPrice, instance.SellPrice, instance.Name, instance.Status, instance.Id)
	if err != nil {
		return ctx, entity.TblInstanceType{}, err
	}

	return ctx, instance, nil
}