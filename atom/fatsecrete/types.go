package fatsecrete

type Food struct {
	FoodName     string
	Brand        string
	Portion      string
	Calorie      string
	Fat          string
	Carbohydrate string
	Protein      string
	Cholestrol   string
	Fiber        string
	Sugar        string
	Sodium       string
	Kalium       string
	Categories   string
}

type InsertByNameReq struct {
	Q string `json:"q"`
}
