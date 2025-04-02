package image_creator

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func GetJpegSource(path string) (image.Image, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("Failed to open image: %w", err)
	}

	defer file.Close()

	src, err := jpeg.Decode(file)

	if err != nil {
		return nil, fmt.Errorf("Failed to decode jpeg: %w", err)
	}

	return src, nil
}
