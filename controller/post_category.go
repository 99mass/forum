package controller

import (
	"database/sql"
)

//Créer un PostCategory en utlisant les deux clés étrangers comme clé primaire
func CreatePostCategory(db *sql.DB, postID, categoryID string) error {
	query := `
        INSERT INTO posts_categories (post_id, category_id)
        VALUES (?, ?);
    `

	_, err := db.Exec(query, postID, categoryID)
	if err != nil {
		return err
	}

	return nil
}

//supprimer un postCategory en utilisant les IDs de post et de category
func DeletePostCategory(db *sql.DB, postID, categoryID string) error {
	query := `
        DELETE FROM posts_categories
        WHERE post_id = ? AND category_id = ?;
    `

	_, err := db.Exec(query, postID, categoryID)
	if err != nil {
		return err
	}

	return nil
}

//mettre à jour un postCategory en utilisant les IDs de post et de category
func UpdatePostCategory(db *sql.DB, postID, categoryID string) error {
	query := `
        UPDATE posts_categories
        SET category_id = ?
        WHERE post_id = ?;
    `

	_, err := db.Exec(query, categoryID, postID)
	if err != nil {
		return err
	}

	return nil
}
