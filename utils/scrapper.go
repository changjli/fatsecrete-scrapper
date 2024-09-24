package scrapper

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type ScrapperInput struct {
	Query string
	Pages int
}

type Product struct {
	Name     string
	Price    string
	Rating   string
	Sales    string
	Shop     string
	Location string
}

func TokopediaScrapper(input ScrapperInput) {
	timestart := time.Now()

	// selenium
	// initialize a Chrome browser instance on port 4444
	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)
	if err != nil {
		log.Fatal("Error:", err)
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
		log.Fatal("Error:", err)
	}

	// maximize the current window to avoid responsive rendering
	err = driver.MaximizeWindow("")
	if err != nil {
		log.Fatal("Error:", err)
	}

	var doneScrapping = make(chan bool)
	var scrappedData = make(chan []Product)

	for page := 1; page <= input.Pages; page++ {
		scrapByPage(driver, input.Query, page, doneScrapping, scrappedData)
		// // request page
		// err = driver.Get(fmt.Sprintf("https://www.tokopedia.com/search?q=%s&page=%d", input.Query, page))
		// fmt.Printf("https://www.tokopedia.com/search?q=%s&page=%d", input.Query, page)
		// if err != nil {
		// 	log.Fatal("Error:", err)
		// }

		// // html, err := driver.PageSource()
		// // if err != nil {
		// // 	log.Fatal("Error:", err)
		// // }
		// // fmt.Println(html)

		// productElements, err := driver.FindElements(selenium.ByCSSSelector, ".prd_container-card")
		// if err != nil {
		// 	log.Fatal("Error:", err)
		// }

		// var products []Product
		// for _, productElement := range productElements {
		// 	// nameElement, err := productElement.FindElement(selenium.ByCSSSelector, ".prd_link-product-name")
		// 	// if err != nil {
		// 	// 	log.Fatal("Error:", err)
		// 	// }
		// 	// priceElement, err := productElement.FindElement(selenium.ByCSSSelector, ".prd_link-product-price")
		// 	// if err != nil {
		// 	// 	log.Fatal("Error:", err)
		// 	// }

		// 	// var ratingElement selenium.WebElement
		// 	// var salesElement selenium.WebElement
		// 	// ratingLabelElement, _ := productElement.FindElement(selenium.ByCSSSelector, ".prd_shop-rating-average-and-label")
		// 	// if ratingLabelElement != nil {
		// 	// 	ratingElement, err = ratingLabelElement.FindElement(selenium.ByCSSSelector, ".prd_rating-average-text")
		// 	// 	if err != nil {
		// 	// 		log.Fatal("Error:", err)
		// 	// 	}
		// 	// 	salesElement, err = ratingLabelElement.FindElement(selenium.ByCSSSelector, ".prd_label-integrity")
		// 	// 	if err != nil {
		// 	// 		log.Fatal("Error:", err)
		// 	// 	}

		// 	// 	rating, err := ratingElement.Text()
		// 	// 	if err != nil {
		// 	// 		log.Fatal("Error:", err)
		// 	// 	}
		// 	// 	sales, err := salesElement.Text()
		// 	// 	if err != nil {
		// 	// 		log.Fatal("Error:", err)
		// 	// 	}

		// 	// 	product.Rating = rating
		// 	// 	product.Sales = sales
		// 	// }

		// 	// var shopElement selenium.WebElement
		// 	// var locationElement selenium.WebElement
		// 	// shopLocationElement, _ := productElement.FindElement(selenium.ByCSSSelector, ".css-z1kcla")
		// 	// if shopLocationElement != nil {
		// 	// 	shopElement, err = shopLocationElement.FindElement(selenium.ByCSSSelector, ".prd_link-shop-name")
		// 	// 	if err != nil {
		// 	// 		log.Fatal("Error:", err)
		// 	// 	}
		// 	// 	locationElement, err = shopLocationElement.FindElement(selenium.ByCSSSelector, ".prd_link-shop-loc")
		// 	// 	if err != nil {
		// 	// 		log.Fatal("Error:", err)
		// 	// 	}

		// 	// 	shop, err := shopElement.Text()
		// 	// 	if err != nil {
		// 	// 		log.Fatal("Error:", err)
		// 	// 	}
		// 	// 	location, err := locationElement.Text()
		// 	// 	if err != nil {
		// 	// 		log.Fatal("Error:", err)
		// 	// 	}

		// 	// 	product.Shop = shop
		// 	// 	product.Location = location
		// 	// }

		// 	// add the scraped data to the list
		// 	product := Product{}
		// 	product.Name = getElementByClass(".prd_link-product-name", productElement)
		// 	product.Price = getElementByClass(".prd_link-product-price", productElement)
		// 	product.Rating = getElementByClass("prd_rating-average-text", productElement)
		// 	product.Sales = getElementByClass(".prd_label-integrity", productElement)
		// 	product.Shop = getElementByClass(".prd_link-shop-name", productElement)
		// 	product.Location = getElementByClass(".prd_link-shop-loc", productElement)

		// 	products = append(products, product)
		// }

		// log.Println(products)
		// log.Println(len(products))
	}

	if <-doneScrapping {
		log.Println(scrappedData)
	}

	totalResponseTime := time.Since(timestart).Seconds()
	log.Println(totalResponseTime)
}

func getElementByClass(class string, parent selenium.WebElement) string {
	element, err := parent.FindElement(selenium.ByCSSSelector, class)
	if err != nil {
		return "-"
	}

	value, err := element.Text()
	if err != nil {
		return "-"
	}

	return value
}

func scrapByPage(driver selenium.WebDriver, query string, page int, doneScrapping chan bool, scrappingData chan []Product) []Product {
	// request page
	err := driver.Get(fmt.Sprintf("https://www.tokopedia.com/search?q=%s&page=%d", query, page))
	fmt.Printf("https://www.tokopedia.com/search?q=%s&page=%d", query, page)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// perform the scrolling interaction
	scrollingScript := `
		// scroll down the page 10 times
		const scrolls = 10
		let scrollCount = 0
			
		// scroll down and then wait for 0.5s
		const scrollInterval = setInterval(() => {
			window.scrollTo(0, document.body.scrollHeight)
			scrollCount++
			if (scrollCount === scrolls) {
			clearInterval(scrollInterval)
			}
		}, 500)
   	`
	_, err = driver.ExecuteScript(scrollingScript, []interface{}{})
	if err != nil {
		log.Fatal("Error:", err)
	}

	// wait up to 10 seconds for the 60th product to be on the page
	err = driver.WaitWithTimeout(func(driver selenium.WebDriver) (bool, error) {
		lastProduct, _ := driver.FindElement(selenium.ByCSSSelector, ".post:nth-of-type(60                                 )")
		if lastProduct != nil {
			return lastProduct.IsDisplayed()
		}
		return false, nil
	}, 20*time.Second)
	if err != nil {
		log.Fatal("Error:", err)
	}

	productElements, err := driver.FindElements(selenium.ByCSSSelector, ".prd_container-card")
	if err != nil {
		log.Fatal("Error:", err)
	}

	var products []Product
	for _, productElement := range productElements {
		// add the scraped data to the list
		product := Product{}
		product.Name = getElementByClass(".prd_link-product-name", productElement)
		product.Price = getElementByClass(".prd_link-product-price", productElement)
		product.Rating = getElementByClass("prd_rating-average-text", productElement)
		product.Sales = getElementByClass(".prd_label-integrity", productElement)
		product.Shop = getElementByClass(".prd_link-shop-name", productElement)
		product.Location = getElementByClass(".prd_link-shop-loc", productElement)

		products = append(products, product)
	}

	log.Println(products, len(products))

	doneScrapping <- true
	scrappingData <- products

	return products
}

func FatSecreteScrapper(foodName string) {
	timestart := time.Now()

	// selenium
	// initialize a Chrome browser instance on port 4444
	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)
	if err != nil {
		log.Fatal("Error:", err)
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
		log.Fatal("Error:", err)
	}

	err = driver.Get(fmt.Sprintf("https://www.fatsecret.co.id/kalori-gizi/search?q=%s", foodName))
	if err != nil {
		log.Fatal("Error:", err)
	}

	var links []string
	for {
		// paging
		nextElement, err := driver.FindElement(selenium.ByCSSSelector, ".next")
		if err != nil {
			log.Fatal("Error:", err)
			break
		}

		nextLinkElement, err := nextElement.FindElement(selenium.ByCSSSelector, "a")
		if err != nil {
			log.Fatal("Error:", err)
		}

		next, err := nextLinkElement.GetAttribute("href")
		if err != nil {
			log.Fatal("Error:", err)
		}

		linkElements, err := driver.FindElements(selenium.ByCSSSelector, ".prominent")
		if err != nil {
			log.Fatal("Error:", err)
		}

		for _, linkElement := range linkElements {
			href, err := linkElement.GetAttribute("href")
			if err != nil {
				log.Fatal("Error:", err)
			}

			links = append(links, href)
		}

		err = driver.Get(next)
		if err != nil {
			log.Fatal("Error:", err)
		}
	}

	log.Printf("%d links found", len(links))

	type Food struct {
		FoodName     string
		Brand        string
		Portion      string
		Calorie      string
		Fat          string
		Carbohydrate string
		Protein      string
	}

	var foods []Food

	for _, link := range links {
		var food Food

		err = driver.Get(link)
		if err != nil {
			log.Fatal("Error:", err)
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
				log.Fatal("Error:", err)
			}

			food.Brand = brand
		}

		// food
		foodNameElement, err := driver.FindElement(selenium.ByCSSSelector, "h1")
		if err != nil {
			log.Fatal("Error:", err)
		}

		foodName, err := foodNameElement.Text()
		if err != nil {
			log.Fatal("Error:", err)
		}

		food.FoodName = foodName

		// nutrients
		factElements, err := driver.FindElements(selenium.ByCSSSelector, ".fact")
		if err != nil {
			log.Fatal("Error:", err)
		}

		for _, factElement := range factElements {

			factTitleElement, err := factElement.FindElement(selenium.ByCSSSelector, ".factTitle")
			if err != nil {
				log.Fatal("Error:", err)
			}

			factValueElement, err := factElement.FindElement(selenium.ByCSSSelector, ".factValue")
			if err != nil {
				log.Fatal("Error:", err)
			}

			factTitle, err := factTitleElement.Text()
			if err != nil {
				log.Fatal("Error:", err)
			}

			factValue, err := factValueElement.Text()
			if err != nil {
				log.Fatal("Error:", err)
			}

			if factTitle == "Kal" {
				food.Calorie = factValue
			} else if factTitle == "Lemak" {
				food.Fat = factValue
			} else if factTitle == "Karb" {
				food.Carbohydrate = factValue
			} else if factTitle == "Prot" {
				food.Protein = factValue
			}
		}

		foods = append(foods, food)
	}

	log.Println(foods)

	totalResponseTime := time.Since(timestart).Seconds()
	log.Println(totalResponseTime)
}
