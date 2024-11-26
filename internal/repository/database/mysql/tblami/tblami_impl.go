package tblami

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"
	"database/sql"
)

type TblAmiImpl struct {
	DB *sql.DB
}

func New(DB *sql.DB) database.TblAmi {
	return &TblAmiImpl{
		DB: DB,
	}
}

func (database *TblAmiImpl) GetAmi(ctx context.Context, key, value string) (context.Context, entity.TblAmi, error) {
	SQL := `
		SELECT 
			id,
			name,
			image_id,
			minecraft_edition,
			minecraft_version,
			mod_loader,
			mod_loader_version,
			minimum_ram,
			minimum_vcpu
		FROM 
			tbl_ami
		WHERE 
			` + key + ` = ?
	`

	row := database.DB.QueryRowContext(ctx, SQL, value)
	var ami entity.TblAmi

	err := row.Scan(
		&ami.Id,
		&ami.Name,
		&ami.ImageId,
		&ami.MinecraftEdition,
		&ami.MinecraftVersion,
		&ami.ModLoader,
		&ami.ModLoaderVersion,
		&ami.MinimumRam,
		&ami.MinimumVcpu,
	)

	if err != nil {
		return ctx, entity.TblAmi{}, err
	}

	return nil, ami, nil
}

func (database *TblAmiImpl) GetAllAmi(ctx context.Context) (context.Context, []entity.TblAmi, error) {
	SQL := `
		SELECT
			id,
			name,
			image_id,
			minecraft_edition,
			minecraft_version,
			mod_loader,
			mod_loader_version,
			minimum_ram,
			minimum_vcpu
		FROM 
			tbl_ami
	`

	rows, err := database.DB.QueryContext(ctx, SQL)
	if err != nil {
		return ctx, nil, err
	}
	defer rows.Close()

	var amis []entity.TblAmi
	for rows.Next() {
		ami := entity.TblAmi{}
		err := rows.Scan(
			&ami.Id,
			&ami.Name,
			&ami.ImageId,
			&ami.MinecraftEdition,
			&ami.MinecraftVersion,
			&ami.ModLoader,
			&ami.ModLoaderVersion,
			&ami.MinimumRam,
			&ami.MinimumVcpu,
		)
		if err != nil {
			return ctx, nil, err
		}
		amis = append(amis, ami)
	}

	return ctx, amis, nil
}
