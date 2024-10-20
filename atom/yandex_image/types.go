package yandex_image

type YandexImage struct {
	Url    string  `json:"url"`
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
}

type MasterFood struct {
	Id    int
	Name  string
	Brand string
}
