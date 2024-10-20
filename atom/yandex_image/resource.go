package yandex_image

import (
	"log"
	db_connection "scrapper/config/postgres"
)

func InsertAllService() error {
	getMasterFoods, err := GetAllMasterFoodRepo()
	if err != nil {
		log.Println("[yandex-image][resource][InsertAllService]:", err)
		return err
	}

	// var dummies []MasterFood = []MasterFood{
	// 	MasterFood{
	// 		Id:    1,
	// 		Name:  "ketoprak",
	// 		Brand: "umum",
	// 	},
	// }

	for _, masterFood := range getMasterFoods {
		getUrls, err := YandexImageScrapper(masterFood)
		if err != nil {
			log.Println("[yandex-image][resource][InsertAllService]:", err)
			return err
		}

		for _, url := range getUrls {
			getFile, err := YandexImageDownload(url, masterFood, "jpg")
			if err != nil {
				log.Println("[yandex-image][resource][InsertAllService]:", err)
				continue
			} else {
				log.Println(getFile)
				break
			}
		}
	}

	return nil
}

func InsertAllServiceV2() error {
	db := db_connection.OpenConnection()
	defer db.Close()

	// getMasterFoods, err := GetAllMasterFoodRepo()
	// if err != nil {
	// 	log.Println("[yandex-image][resource][InsertAllService]:", err)
	// 	return err
	// }

	var dummies []MasterFood = []MasterFood{
		{
			Id:    1,
			Name:  "ketoprak",
			Brand: "umum",
		},
		{
			Id:    1,
			Name:  "ketoprak salad bowl",
			Brand: "salad point",
		},
		{
			Id:    1,
			Name:  "telur rebus",
			Brand: "umum",
		},
		{
			Id:    1,
			Name:  "telur",
			Brand: "umum",
		},
		{
			Id:    1,
			Name:  "telur dadar",
			Brand: "umum",
		},
	}

	getYandexImages, err := YandexImageScrapperV2(dummies)
	if err != nil {
		log.Println("[yandex-image][resource][InsertAllService]:", err)
		return err
	}

	for _, yandexImage := range getYandexImages {
		for _, url := range yandexImage.Urls {
			getFile, err := YandexImageDownload(url, yandexImage.MasterFood, "jpg")
			if err != nil {
				log.Println("[yandex-image][resource][InsertAllService]:", err)
				continue
			} else {
				yandexImage.Url = url
				yandexImage.File = getFile
				err := InsertImageRepo(db, yandexImage)
				if err != nil {
					log.Println("[yandex-image][resource][InsertAllService]:", err)
					return err
				}
				break
			}
		}

	}

	return nil
}
