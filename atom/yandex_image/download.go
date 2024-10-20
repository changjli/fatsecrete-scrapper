package yandex_image

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func YandexImageDownload(url string, masterFood MasterFood, format string) (string, error) {
	// Download the image
	resp, err := http.Get(url)
	if err != nil {
		err := errors.New("fetch error")
		log.Println("[yandex-image][download][YandexImageDownload]", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("fetch image failed with status code %d", resp.StatusCode)
		log.Println("[yandex-image][download][YandexImageDownload]", err)
		return "", err
	}

	contentType := resp.Header.Get("Content-Type")
	fmt.Println("Content-Type:", contentType)

	// Save the image as jpg or png
	fileName := Slugify(fmt.Sprintf("%s %s", masterFood.Name, masterFood.Brand))
	outputFile, err := os.Create(fmt.Sprintf("storage/%s.%s", fileName, "jpeg"))
	if err != nil {
		err := errors.New("error downloading file")
		log.Println("[yandex-image][download][YandexImageDownload]", err)
		return "", err
	}
	defer outputFile.Close()

	// Copy the response body directly to the file
	_, err = io.Copy(outputFile, resp.Body)
	if err != nil {
		log.Println("[yandex-image][download][YandexImageDownload]: failed to write to file", err)
		return "", err
	}

	return fileName, nil
}

func Slugify(keyword string) string {
	strArr := strings.Split(strings.ToLower(keyword), " ")
	return strings.Join(strArr, "-")
}
