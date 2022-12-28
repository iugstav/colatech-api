package imgmock

import (
	"fmt"
	"image"
	"image/color"
)

func CreateMockImage() (image.Image, error) {
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if (x+y)%2 == 0 {
				img.Set(x, y, color.White)
			} else {
				img.Set(x, y, color.Black)
			}
		}
	}
	return img, nil
}

func CompareImages(img1, img2 image.Image) error {
	if img1.Bounds() != img2.Bounds() {
		return fmt.Errorf("image dimensions do not match")
	}
	for y := 0; y < img1.Bounds().Dy(); y++ {
		for x := 0; x < img1.Bounds().Dx(); x++ {
			c1 := img1.At(x, y)
			c2 := img2.At(x, y)
			if c1 != c2 {
				return fmt.Errorf("pixels at (%d, %d) do not match: %v != %v", x, y, c1, c2)
			}
		}
	}
	return nil
}
