package tblparameters

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"
	"database/sql"
)

type TblParametersImpl struct {
	DB *sql.DB
}

func New(DB *sql.DB) database.TblParameters {
	return &TblParametersImpl{
		DB: DB,
	}
}

func (database *TblParametersImpl) GetAllParameters(ctx context.Context) (context.Context, []entity.TblParameters, error) {
	query := "SELECT name, value, description FROM tbl_parameters"
	rows, err := database.DB.QueryContext(ctx, query)
	if err != nil {
		return ctx, nil, err
	}
	defer rows.Close()

	var parameters []entity.TblParameters
	for rows.Next() {
		var param entity.TblParameters
		if err := rows.Scan(&param.Name, &param.Value, &param.Description); err != nil {
			return ctx, nil, err
		}
		parameters = append(parameters, param)
	}

	if err := rows.Err(); err != nil {
		return ctx, nil, err
	}

	return ctx, parameters, nil
}

func (database *TblParametersImpl) GetParameters(ctx context.Context, key, value string) (context.Context, entity.TblParameters, error) {
	query := "SELECT name, value, description FROM tbl_parameters WHERE Name = ? AND Value = ?"
	row := database.DB.QueryRowContext(ctx, query, key, value)

	var param entity.TblParameters
	if err := row.Scan(&param.Name, &param.Value, &param.Description); err != nil {
		if err == sql.ErrNoRows {
			return ctx, entity.TblParameters{}, nil
		}
		return ctx, entity.TblParameters{}, err
	}

	return ctx, param, nil
}