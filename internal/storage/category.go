package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/EwvwGeN/InHouseAd_assignment/internal/domain/models"
	"github.com/jackc/pgconn"
)

func (pp *postgresProvider) SaveCategory(ctx context.Context, category models.Category) error {
	_, err := pp.dbConn.Exec(ctx, fmt.Sprintf(`INSERT INTO "%s" (name, code, description)
VALUES($1,$2,$3);`,
	pp.cfg.CatogoryTable),
	category.Name,
	category.Code,
	category.Description)
	if err == nil {
		return nil
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			return ErrCategoryExist
		}
	}
	return ErrQuery
}

func (pp *postgresProvider) GetCategoryByCode(ctx context.Context, catCode string) (models.Category, error) {
	row := pp.dbConn.QueryRow(ctx, fmt.Sprintf(`
SELECT "name", "code", "description"
FROM "%s"
WHERE "code"=$1;`,
	pp.cfg.CatogoryTable),
	catCode)
	var (
		category models.Category
	)
	err := row.Scan(&category.Name, &category.Code, &category.Description)
	if err != nil {
		return models.Category{}, ErrQuery
	}
	return category, nil
}

func (pp *postgresProvider) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	rows, err := pp.dbConn.Query(ctx, fmt.Sprintf(`
SELECT "name", "code", "description"
FROM "%s"`,
	pp.cfg.CatogoryTable))
	if err != nil {
		return nil, ErrQuery
	}
	var outCategorys []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.Name, &category.Code, &category.Description)
		if err != nil {
			return nil, ErrQuery
		}
		outCategorys = append(outCategorys, category)
	}
	return outCategorys, nil
}

func (pp *postgresProvider) UpdateCategoryByCode(ctx context.Context, catCode string, catUpdateData models.CategoryForPatch) error {
	preparedQuery := fmt.Sprintf("UPDATE \"%s\" SET ", pp.cfg.CatogoryTable)
	// is it faster to use marshal to json and unmarshal to map[string]interface{} and then range it by for statement?
	usedFields := 0
	usedData := make([]interface{}, 0)
	if catUpdateData.Name != nil {
		preparedQuery += fmt.Sprintf("\"name\" = $%d, ", usedFields+1)
		usedFields++
		usedData = append(usedData, *catUpdateData.Name)
	}
	if catUpdateData.Code != nil {
		preparedQuery += fmt.Sprintf("\"code\" = $%d, ", usedFields+1)
		usedFields++
		usedData = append(usedData, *catUpdateData.Code)
	}
	if catUpdateData.Description != nil {
		preparedQuery += fmt.Sprintf("\"description\" = $%d, ", usedFields+1)
		usedFields++
		usedData = append(usedData, *catUpdateData.Description)
	}
	// the worst but fast solution
	preparedQuery = preparedQuery[:len(preparedQuery)-2]
	usedData = append(usedData, catCode)
	_, err := pp.dbConn.Exec(ctx, fmt.Sprintf("%s WHERE \"code\" = $%d", preparedQuery, usedFields+1), usedData...)
	if err != nil {
		return ErrQuery
	}
	return nil
}

func (pp *postgresProvider) DeleteCategoryBycode(ctx context.Context, catCode string) error {
	_, err := pp.dbConn.Exec(ctx, fmt.Sprintf("DELETE FROM \"%s\" WHERE \"code\" = $1", pp.cfg.CatogoryTable), catCode)
	if err != nil {
		return ErrQuery
	}
	return nil
}
