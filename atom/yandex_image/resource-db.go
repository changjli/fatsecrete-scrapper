package yandex_image

import (
	"log"
	db_connection "scrapper/config/postgres"
)

func GetAllMasterFoodDB() ([]MasterFood, error) {
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
