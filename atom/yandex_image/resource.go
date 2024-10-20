package yandex_image

import "log"

func InsertAllService() error {
	getMasterFoods, err := GetAllMasterFoodDB()
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
