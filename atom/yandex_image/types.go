package yandex_image

type YandexImage struct {
	MasterFood MasterFood
	Urls       []string
	Url        string
	File       string
}

type MasterFood struct {
	Id    int
	Name  string
	Brand string
}
