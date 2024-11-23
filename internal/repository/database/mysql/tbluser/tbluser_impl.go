package tbluser

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"
	"database/sql"
)

type TblUserImpl struct {
	DB *sql.DB
}

func New(DB *sql.DB) database.TblUser {
	return &TblUserImpl{
		DB: DB,
	}
}

func (database *TblUserImpl) InsertUser(ctx context.Context, user entity.TblUser) (context.Context, entity.TblUser, error) {
	SQL := `
    INSERT INTO tbl_user (username, password)
    VALUES (?, ?)
  `

	row := database.DB.QueryRowContext(ctx, SQL, user.Username, user.Password)
	paramData := entity.TblUser{}
	err := row.Scan(
		&paramData.Id,
		&paramData.Username,
		&paramData.Password,
		&paramData.CreatedDate,
		&paramData.ModifiedDate,
	)
	if err != nil {
		return ctx, paramData, err
	}

	return ctx, paramData, nil
}

func (database *TblUserImpl) GetUser(ctx context.Context, key string, value string) (context.Context, entity.TblUser, error) {
	SQL := `
    SELECT id, username, password, created_date, modified_date
    FROM tbl_user
    WHERE ` + key + ` = ?
  `

	row := database.DB.QueryRowContext(ctx, SQL, value)
	paramData := entity.TblUser{}
	err := row.Scan(
		&paramData.Id,
		&paramData.Username,
		&paramData.Password,
		&paramData.CreatedDate,
		&paramData.ModifiedDate,
	)
	if err != nil {
		return ctx, paramData, err
	}

	return ctx, paramData, nil
}

func (database *TblUserImpl) PaginateUser(ctx context.Context, page int, size int) (context.Context, []entity.TblUser, error) {
	offset := (page - 1) * size
	SQL := `
		SELECT id, username, password, created_date, modified_date
		FROM tbl_user
		LIMIT ? OFFSET ?
	`

	rows, err := database.DB.QueryContext(ctx, SQL, size, offset)
	if err != nil {
		return ctx, nil, err
	}
	defer rows.Close()

	var paramData []entity.TblUser
	for rows.Next() {
		data := entity.TblUser{}
		err := rows.Scan(
			&data.Id,
			&data.Username,
			&data.Password,
			&data.CreatedDate,
			&data.ModifiedDate,
		)
		if err != nil {
			return ctx, nil, err
		}
		paramData = append(paramData, data)
	}

	return ctx, paramData, nil
}
