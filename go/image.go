package openapi

import (
	"errors"
	"fmt"
	"image"
	"os"
)

func decodeJpgImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fi, _ := f.Stat()
	fmt.Println(fi.Name())

	img, format, err := image.Decode(f)
	if err != nil {
		fmt.Println("Decoding error:", err.Error())
		return nil, err
	}
	if format != "jpeg" {
		fmt.Println("image format is not jpeg")
		return nil, errors.New("")
	}
	return img, nil
}
