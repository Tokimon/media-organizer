package image_creator

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"

	"golang.org/x/image/draw"
)

func GenerateThumbnail(writer io.Writer, img image.Image) error {
	return ScaleImage(writer, img, 300)
}

func ScaleImage(writer io.Writer, img image.Image, size int) error {
	bounds := img.Bounds()
	width := float32(bounds.Dx())
	height := float32(bounds.Dy())

	tw := float32(size)
	th := tw * (height / width)

	if width < height {
		th = tw
		tw = th * (width / height)
	}

	thumbnail := image.NewRGBA(image.Rect(0, 0, int(tw), int(th)))
	draw.NearestNeighbor.Scale(thumbnail, thumbnail.Rect, img, bounds, draw.Over, nil)

	err := jpeg.Encode(writer, thumbnail, nil)

	if err != nil {
		return fmt.Errorf("Failed to encode thumbnail: %w", err)
	}

	return nil
}
