package desktop

import (
	"image"

	"github.com/kbinani/screenshot"
)

func GetScreenshots() ([]image.Image, error) {
	screenshots := []image.Image{}
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			return nil, err
		}
		screenshots = append(screenshots, img)
	}
	return screenshots, nil
}
