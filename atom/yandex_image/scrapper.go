package yandex_image

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func YandexImageScrapper(masterFood MasterFood) ([]string, error) {
	// selenium
	// initialize a Chrome browser instance on port 4444
	service, err := selenium.NewChromeDriverService("./chromedriver.exe", 4444)
	if err != nil {
		log.Println("[initializeSelenium]:", err)
		return nil, err
	}
	defer service.Stop()

	// configure the browser options
	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless-new", // comment out this line for testing
	}})

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Println("[initializeSelenium]:", err)
		return nil, err
	}

	err = driver.Get(fmt.Sprintf("https://yandex.com/images/search?text=%s", fmt.Sprintf("%s %s", masterFood.Name, masterFood.Brand)))
	if err != nil {
		log.Println("[yandex-image][scrapper][YandexImageScrapper]:", err)
		return nil, err
	}

	time.Sleep(1 * time.Second)

	imgElements, err := driver.FindElements(selenium.ByXPATH, "//*[contains(@class, 'ContentImage-Image') and contains(@class, 'ContentImage-Image_clickable')]")
	if err != nil {
		log.Println("[yandex-image][scrapper][YandexImageScrapper]:", err)
		return nil, err
	}

	var srcs []string
	for _, imgElement := range imgElements {
		// get the natural width and height of the image
		widthStr, err := imgElement.GetAttribute("naturalWidth")
		if err != nil {
			log.Println("[yandex-image][scrapper][YandexImageScrapper]:", err)
			continue
		}

		heightStr, err := imgElement.GetAttribute("naturalHeight")
		if err != nil {
			log.Println("[yandex-image][scrapper][YandexImageScrapper]:", err)
			continue
		}

		// convert the width and height from string to int
		width, err := strconv.Atoi(widthStr)
		if err != nil {
			log.Println("[yandex-image][scrapper][YandexImageScrapper]:", err)
			continue
		}

		height, err := strconv.Atoi(heightStr)
		if err != nil {
			log.Println("[yandex-image][scrapper][YandexImageScrapper]:", err)
			continue
		}

		// Ensure potrait images
		if width > height {
			src, err := imgElement.GetAttribute("src")
			if err != nil {
				log.Println("[yandex-image][scrapper][YandexImageScrapper]:", err)
				return nil, err
			}
			srcs = append(srcs, src)
		}
	}

	return srcs, nil
}

func YandexImageScrapperV2(masterFoods []MasterFood) ([]YandexImage, error) {
	// selenium
	// initialize a Chrome browser instance on port 4444
	service, err := selenium.NewChromeDriverService("./chromedriver.exe", 4444)
	if err != nil {
		log.Println("[initializeSelenium]:", err)
		return nil, err
	}
	defer service.Stop()

	// configure the browser options
	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless-new", // comment out this line for testing
	}})

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Println("[initializeSelenium]:", err)
		return nil, err
	}

	var datas []YandexImage
	for _, masterFood := range masterFoods {
		var query string
		if masterFood.Brand == "umum" {
			query = masterFood.Name
		} else {
			query = fmt.Sprintf("%s %s", masterFood.Name, masterFood.Brand)
		}

		err = driver.Get(fmt.Sprintf("https://yandex.com/images/search?text=%s", query))
		if err != nil {
			log.Println("[yandex-image][scrapper][YandexImageScrapper]:", err)
			return nil, err
		}

		time.Sleep(1 * time.Second)

		imgElements, err := driver.FindElements(selenium.ByXPATH, "//*[contains(@class, 'ContentImage-Image') and contains(@class, 'ContentImage-Image_clickable')]")
		if err != nil {
			log.Println("[yandex-image][scrapper][YandexImageScrapper]:", err)
			return nil, err
		}

		var srcs []string
		for _, imgElement := range imgElements {
			// get the natural width and height of the image
			widthStr, err := imgElement.GetAttribute("naturalWidth")
			if err != nil {
				log.Println("[yandex-image][scrapper][YandexImageScrapper]:", err)
				continue
			}

			heightStr, err := imgElement.GetAttribute("naturalHeight")
			if err != nil {
				log.Println("[yandex-image][scrapper][YandexImageScrapper]:", err)
				continue
			}

			// convert the width and height from string to int
			width, err := strconv.Atoi(widthStr)
			if err != nil {
				log.Println("[yandex-image][scrapper][YandexImageScrapper]:", err)
				continue
			}

			height, err := strconv.Atoi(heightStr)
			if err != nil {
				log.Println("[yandex-image][scrapper][YandexImageScrapper]:", err)
				continue
			}

			// Ensure potrait images
			if width > height {
				src, err := imgElement.GetAttribute("src")
				if err != nil {
					log.Println("[yandex-image][scrapper][YandexImageScrapper]:", err)
					return nil, err
				}
				srcs = append(srcs, src)
			}
		}

		datas = append(datas, YandexImage{
			MasterFood: masterFood,
			Urls:       srcs,
		})
	}

	return datas, nil
}
