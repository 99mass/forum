package controller

import (
	"database/sql"
	"errors"

	"forum/models"
)

func CreateCategory(db *sql.DB, category models.Category) (int64, error) {
	query := `
        INSERT INTO categories (name_category)
        VALUES (?);
    `

	result, err := db.Exec(query, category.NameCategory)
	if err != nil {
		return 0, err
	}

	categoryID, _ := result.LastInsertId()
	return categoryID, nil
}

func GetAllCategories(db *sql.DB) ([]models.Category, error) {
	query := `
        SELECT id, name_category
        FROM categories;
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.NameCategory)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func GetCategoryByID(db *sql.DB, categoryID int64) (models.Category, error) {
	var category models.Category
	query := `
        SELECT id, name_category
        FROM categories
        WHERE id = ?
        LIMIT 1;
    `

	err := db.QueryRow(query, categoryID).Scan(&category.ID, &category.NameCategory)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Category{}, errors.New("catégorie non trouvée")
		}
		return models.Category{}, err
	}

	return category, nil
}

func UpdateCategory(db *sql.DB, category models.Category) error {
	query := `
        UPDATE categories
        SET name_category = ?
        WHERE id = ?;
    `

	_, err := db.Exec(query, category.NameCategory, category.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCategory(db *sql.DB, categoryID int64) error {
	query := `
        DELETE FROM categories
        WHERE id = ?;
    `

	_, err := db.Exec(query, categoryID)
	if err != nil {
		return err
	}

	return nil
}
