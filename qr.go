package frothy

import (
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
