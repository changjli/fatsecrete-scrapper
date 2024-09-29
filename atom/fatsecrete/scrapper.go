package fatsecrete

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func FatSecreteScrapper(foodName string) ([]Food, error) {
	timestart := time.Now()

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

	err = driver.Get(fmt.Sprintf("https://www.fatsecret.co.id/kalori-gizi/search?q=%s", foodName))
	if err != nil {
		log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
		return nil, err
	}

	var links []string
	for {
		// prevent page not working
		time.Sleep(2 * time.Second)

		// links
		linkElements, err := driver.FindElements(selenium.ByCSSSelector, ".prominent")
		if err != nil {
			log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
			return nil, err
		}

		for _, linkElement := range linkElements {
			href, err := linkElement.GetAttribute("href")
			if err != nil {
				log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
				return nil, err
			}

			links = append(links, href)
		}

		// paging
		nextElement, err := driver.FindElement(selenium.ByCSSSelector, ".next")
		if err != nil {
			break
		}

		nextLinkElement, err := nextElement.FindElement(selenium.ByCSSSelector, "a")
		if err != nil {
			log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
			return nil, err
		}

		next, err := nextLinkElement.GetAttribute("href")
		if err != nil {
			log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
			return nil, err
		}

		err = driver.Get(next)
		if err != nil {
			log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
			return nil, err
		}
	}

	log.Printf("%d links found", len(links))

	var foods []Food
	for _, link := range links {
		food, err := FatSecreteScrapPage(driver, link)
		if err != nil {
			log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
			return nil, err
		}
		foods = append(foods, food)
	}

	totalResponseTime := time.Since(timestart).Seconds()
	log.Printf("Scrapping time: %f", totalResponseTime)

	return foods, nil
}

func FatSecreteScrapPage(driver selenium.WebDriver, link string) (Food, error) {
	var food Food

	err := driver.Get(link)
	if err != nil {
		log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
		return Food{}, err
	}

	// prevent page not working
	time.Sleep(2 * time.Second)

	// brand
	brandElement, err := driver.FindElement(selenium.ByCSSSelector, ".manufacturer")
	if err != nil {
		food.Brand = "umum"
	} else {
		brand, err := brandElement.Text()
		if err != nil {
			log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
			return Food{}, err
		}

		food.Brand = brand
	}

	// food
	foodNameElement, err := driver.FindElement(selenium.ByCSSSelector, "h1")
	if err != nil {
		log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
		return Food{}, err
	}

	foodName, err := foodNameElement.Text()
	if err != nil {
		log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
		return Food{}, err
	}

	food.FoodName = foodName

	// portion
	portionElement, err := driver.FindElement(selenium.ByCSSSelector, ".serving_size_value")
	if err != nil {
		log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
		return Food{}, err
	}

	portion, err := portionElement.Text()
	if err != nil {
		log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
		return Food{}, err
	}

	food.Portion = portion

	nutrientFactElement, err := driver.FindElement(selenium.ByCSSSelector, ".nutrition_facts")
	if err != nil {
		log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
		return Food{}, err
	}

	nutrientLabelElements, err := nutrientFactElement.FindElements(selenium.ByCSSSelector, ".nutrient.left")
	if err != nil {
		log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
		return Food{}, err
	}

	nutrientValueElements, err := nutrientFactElement.FindElements(selenium.ByCSSSelector, ".nutrient.right")
	if err != nil {
		log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
		return Food{}, err
	}

	// nutrients v2
	for i, nutrientLabelElement := range nutrientLabelElements {
		if i == 1 {
			continue
		}

		nutrientLabel, err := nutrientLabelElement.Text()
		if err != nil {
			log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
			return Food{}, err
		}

		nutrientValue, err := nutrientValueElements[i+1].Text()
		if err != nil {
			log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
			return Food{}, err
		}

		if nutrientLabel == "Energi" {
			food.Calorie = nutrientValue
		} else if nutrientLabel == "Lemak" {
			food.Fat = nutrientValue
		} else if nutrientLabel == "Kolesterol" {
			food.Cholestrol = nutrientValue
		} else if nutrientLabel == "Protein" {
			food.Protein = nutrientValue
		} else if nutrientLabel == "Karbohidrat" {
			food.Carbohydrate = nutrientValue
		} else if nutrientLabel == "Serat" {
			food.Fiber = nutrientValue
		} else if nutrientLabel == "Gula" {
			food.Sugar = nutrientValue
		} else if nutrientLabel == "Sodium" {
			food.Sodium = nutrientValue
		} else if nutrientLabel == "Kalium" {
			food.Kalium = nutrientValue
		}
	}

	// categories
	detailElement, err := driver.FindElement(selenium.ByCSSSelector, ".details")
	if err != nil {
		log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
		return Food{}, err
	}

	categories := ""
	h4Elements, err := detailElement.FindElements(selenium.ByCSSSelector, "h4")
	if err != nil {
		log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
		return Food{}, err
	}

	for _, h4Element := range h4Elements {
		h4, err := h4Element.Text()
		if err != nil {
			log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
			return Food{}, err
		}

		prefix := "Jenis"

		if strings.HasPrefix(h4, prefix) {
			words := strings.Split(h4, " ")
			category := strings.ReplaceAll(words[len(words)-1], ":", "")
			if categories == "" {
				categories += category
			} else {
				categories += fmt.Sprintf(", %s", category)
			}
		}
	}

	food.Categories = categories

	// // nutrients
	// factElements, err := driver.FindElements(selenium.ByCSSSelector, ".fact")
	// if err != nil {
	// 	log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
	// 	return nil, err
	// }

	// for _, factElement := range factElements {

	// 	factTitleElement, err := factElement.FindElement(selenium.ByCSSSelector, ".factTitle")
	// 	if err != nil {
	// 		log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
	// 		return nil, err
	// 	}

	// 	factValueElement, err := factElement.FindElement(selenium.ByCSSSelector, ".factValue")
	// 	if err != nil {
	// 		log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
	// 		return nil, err
	// 	}

	// 	factTitle, err := factTitleElement.Text()
	// 	if err != nil {
	// 		log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
	// 		return nil, err
	// 	}

	// 	factValue, err := factValueElement.Text()
	// 	if err != nil {
	// 		log.Println("[fatsecrete][scrapper][FatSecreteScrapper]:", err)
	// 		return nil, err
	// 	}

	// 	if factTitle == "Kal" {
	// 		food.Calorie = factValue
	// 	} else if factTitle == "Lemak" {
	// 		food.Fat = factValue
	// 	} else if factTitle == "Karb" {
	// 		food.Carbohydrate = factValue
	// 	} else if factTitle == "Prot" {
	// 		food.Protein = factValue
	// 	}
	// }

	return food, nil
}
