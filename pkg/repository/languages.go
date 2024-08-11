package repository

import (
	"context"

	"github.com/AdelYarR/Data-Blossom/pkg/models"
)

// Create
func (repo *PGRepo) NewLanguage(item models.Languages) error {
	_, err := repo.pool.Exec(context.Background(), `
		INSERT INTO languages (name, type_id) 
		VALUES ($1, $2);`,
		item.Name,
		item.TypeID,
	)
	if err != nil {
		return err
	}

	return nil
}

// Read
func (repo *PGRepo) GetLanguages() ([]models.Languages, error) {
	rows, err := repo.pool.Query(context.Background(),`
	SELECT id, name, type_id FROM languages;`,
	)	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []models.Languages
	for rows.Next() {
		var model models.Languages
		err = rows.Scan(
			&model.ID,
			&model.Name,
			&model.TypeID,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, model)
	}

	return data, nil

}

// Update
func (repo *PGRepo) UpdateLanguage(id string, lang models.Languages) error {
	_, err := repo.pool.Exec(context.Background(), `
		UPDATE languages SET name = $1, type_id = $2
		WHERE id = $3`,
		lang.Name,
		lang.TypeID,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete
func (repo *PGRepo) DeleteLanguage(id string) error {
	_, err := repo.pool.Exec(context.Background(), `
		DELETE FROM languages WHERE id = $1;`,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}