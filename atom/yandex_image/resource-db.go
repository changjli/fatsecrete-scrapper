package yandex_image

import (
	"database/sql"
	"log"
	db_connection "scrapper/config/postgres"
)

func GetAllMasterFoodRepo() ([]MasterFood, error) {
	db := db_connection.OpenConnection()
	defer db.Close()

	query := `
		SELECT
		mfc."id", 
		mfc."food_name", 
		mfc."brand"
		FROM master_foods_combined mfc 
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Println("[yandex-image][resource-db][GetAllMasterFoodDB]", err)
		return nil, err
	}

	var datas []MasterFood
	for rows.Next() {
		var data MasterFood
		rows.Scan(
			&data.Id,
			&data.Name,
			&data.Brand,
		)

		datas = append(datas, data)
	}

	return datas, nil
}

func InsertImageRepo(db *sql.DB, image YandexImage) error {
	query := `
		INSERT INTO master_food_images
		("food_id", "food_name", "food_brand", "url", "file")
		VALUES
		($1, $2, $3, $4, $5)
		ON CONFLICT ("food_name", "food_brand")
		DO UPDATE SET "url" = $4, 
		"file" = $5
	`

	_, err := db.Exec(query, image.MasterFood.Id, image.MasterFood.Name, image.MasterFood.Brand, image.Url, image.File)
	if err != nil {
		log.Println("[yandex-image][resource-db][InsertImageRepo]", err)
		return nil
	}

	return nil
}
