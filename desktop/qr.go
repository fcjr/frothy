package desktop

import (
	"fmt"

	"github.com/fcjr/frothy"
)

func DecodeFromScreen() (string, error) {
	screenshots, err := GetScreenshots()
	if err != nil {
		return "", err
	}
	if len(screenshots) < 1 {
		return "", fmt.Errorf("failed to capture screen")
	}

	var res string
	for _, img := range screenshots {
		res, err = frothy.DecodeQRFromImage(img)
		if err == nil {
			return res, nil
		}
	}
	return "", err
}
