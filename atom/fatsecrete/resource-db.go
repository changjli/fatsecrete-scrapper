package fatsecrete

import (
	"log"
	db_connection "scrapper/config/postgres"
)

func InsertByNameRepo(datas []Food) error {
	db := db_connection.OpenConnection()
	defer db.Close()

	for _, data := range datas {
		queryString :=
			`
			INSERT INTO master_foods (
				food_name, 
				calorie, 
				fat, 
				carbohydrate, 
				protein, 
				brand, 
				src, 
				portion, 
				cholestrol, 
				fiber, 
				sugar, 
				sodium, 
				kalium, 
				categories
			)
			VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
			ON CONFLICT (food_name, brand)
			DO NOTHING;
		`
		_, err := db.Exec(queryString, data.FoodName, data.Calorie, data.Fat, data.Carbohydrate, data.Protein, data.Brand, "fatsecrete", data.Portion,
			data.Cholestrol, data.Fiber, data.Sugar, data.Sodium, data.Kalium, data.Categories)
		if err != nil {
			log.Println("[fatsecrete][resource-db][InsertByNameRepo]", err)
			return err
		}
	}

	return nil
}
