package frothy

import (
	"fmt"
	"image"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func DecodeQRFromImage(img image.Image) (string, error) {
	// prepare BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return "", err
	}

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	res, err := qrReader.Decode(bmp, nil)
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

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
		res, err = DecodeQRFromImage(img)
		if err == nil {
			return res, nil
		}
	}
	return "", err
}
