package fatsecrete

import "log"

func InsertByNameService(req InsertByNameReq) error {
	foods, err := FatSecreteScrapper(req.Q)
	if err != nil {
		log.Println("[fatsecrete][resource][InsertByNameService]:", err)
		return err
	}

	log.Println(foods)
	log.Println("Inserting to database")

	err = InsertByNameRepo(foods)
	if err != nil {
		log.Println("[fatsecrete][resource][InsertByNameService]:", err)
		return err
	}

	return nil
}
