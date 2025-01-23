package services

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"net/http"
	"time"
)

func DownloadImage(url string) (image.Image, int, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error downloading image: %v", err)
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error decoding image: %v", err)
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	return img, width, height, nil
}

func CalculatePerimeter(width, height int) int {
	return 2 * (width + height)
}

func IntroduceRandomDelay() {
	delay := time.Duration(rand.Intn(300)+100) * time.Millisecond
	time.Sleep(delay)
}

func ProcessImages(imageUrls []string) ([]int, error) {
	var results []int

	rand.Seed(time.Now().UnixNano())

	for _, url := range imageUrls {
		_, width, height, err := DownloadImage(url)
		if err != nil {
			return nil, fmt.Errorf("failed to download image %s: %v", url, err)
		}

		perimeter := CalculatePerimeter(width, height)
		results = append(results, perimeter)

		IntroduceRandomDelay()
	}

	return results, nil
}
