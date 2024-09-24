package main

import (
	"runtime"
	"scrapper/routes"
)

func main() {
	runtime.GOMAXPROCS(2)

	// scrapper.TokopediaScrapper(scrapper.ScrapperInput{Query: "gelas%20plastik", Pages: 2})

	// indonesianFoods := []string{
	// 	"Nasi Goreng",   // Fried rice
	// 	"Sate Ayam",     // Chicken satay
	// 	"Gado-Gado",     // Vegetable salad with peanut sauce
	// 	"Rendang",       // Spicy beef stew
	// 	"Nasi Uduk",     // Coconut milk rice
	// 	"Nasi Padang",   // Padang-style rice with various sides
	// 	"Bakso",         // Meatball soup
	// 	"Soto Ayam",     // Chicken soup with turmeric broth
	// 	"Soto Betawi",   // Betawi-style beef soup with coconut milk
	// 	"Pecel Lele",    // Fried catfish with sambal
	// 	"Nasi Kuning",   // Yellow rice
	// 	"Tempe Goreng",  // Fried tempeh
	// 	"Perkedel",      // Fried potato patties
	// 	"Martabak",      // Stuffed pancake, savory or sweet
	// 	"Kerupuk",       // Crackers (shrimp, fish, etc.)
	// 	"Mie Goreng",    // Fried noodles
	// 	"Mie Ayam",      // Chicken noodles
	// 	"Rawon",         // Dark beef soup with black keluak
	// 	"Gudeg",         // Sweet young jackfruit stew
	// 	"Pempek",        // Fish cake from Palembang
	// 	"Capcay",        // Stir-fried mixed vegetables
	// 	"Es Campur",     // Mixed ice dessert with fruits and toppings
	// 	"Sayur Lodeh",   // Vegetable soup with coconut milk
	// 	"Es Cendol",     // Green rice flour jelly dessert
	// 	"Opor Ayam",     // Chicken cooked in coconut milk
	// 	"Nasi Campur",   // Mixed rice with assorted side dishes
	// 	"Siomay",        // Steamed fish dumplings with peanut sauce
	// 	"Serabi",        // Sweet pancake with various toppings
	// 	"Bubur Ayam",    // Chicken porridge
	// 	"Ketoprak",      // Vegetables, tofu, and rice cake with peanut sauce
	// 	"Gudeg",         // Sweet jackfruit stew
	// 	"Lumpia",        // Spring rolls
	// 	"Roti Bakar",    // Toasted bread with various fillings
	// 	"Tahu Gejrot",   // Fried tofu in sweet soy sauce
	// 	"Lontong Sayur", // Rice cake in vegetable curry
	// 	"Bebek Goreng",  // Fried duck
	// 	"Ayam Bakar",    // Grilled chicken
	// 	"Klepon",        // Sweet rice cake balls filled with palm sugar
	// 	"Kue Lapis",     // Layer cake
	// }

	// for _, indonesianFood := range indonesianFoods {
	// 	scrapper.FatSecreteScrapper(indonesianFood)
	// }

	// fatsecrete.FatSecreteScrapper("ketoprak")

	router := routes.SetupRoutes()

	router.Run(":8015")
}
