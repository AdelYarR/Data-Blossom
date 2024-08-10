package repository

import (
	"context"

	"github.com/AdelYarR/Data-Blossom/pkg/models"
)

// Create
func (repo *PGRepo) NewLanguage(item models.Languages) (models.Languages, error) {
	var newLanguage models.Languages
	err := repo.pool.QueryRow(context.Background(), `
		INSERT INTO languages (name, type_id) 
		VALUES ($1, $2) 
		RETURNING id, name, type_id;`,
		item.Name,
		item.TypeID,
	).Scan(&newLanguage.ID, &newLanguage.Name, &newLanguage.TypeID)	
	if err != nil {
		return models.Languages{}, err
	}

	return newLanguage, nil
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
// func (repo *PGRepo) UpdateLanguageById() (models.Languages, error) {
	
// }

// Delete
func (repo *PGRepo) DeleteLanguageById(item models.Languages) (models.Languages, error) {
	var newLanguage models.Languages
	err := repo.pool.QueryRow(context.Background(), `
		DELETE FROM languages WHERE name = $1 && type_id = $2 
		RETURNING id, name, type_id;`,
		item.Name,
		item.TypeID,
	).Scan(&newLanguage.ID, &newLanguage.Name, &newLanguage.TypeID)	
	if err != nil {
		return models.Languages{}, err
	}

	return newLanguage, nil
}