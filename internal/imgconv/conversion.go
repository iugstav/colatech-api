package imgconv

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"net/http"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"github.com/pkg/errors"
)

func ToWebp(image []byte) ([]byte, error) {
	imgType := http.DetectContentType(image)

	switch imgType {
	case "image/webp":
		return image, nil
	case "image/png":
		img, err := png.Decode(bytes.NewReader(image))
		if err != nil {
			return nil, errors.Wrap(err, "unable to decode png")
		}

		webpOptions, optionsErr := encoder.NewLosslessEncoderOptions(encoder.PresetDefault, 75)
		if optionsErr != nil {
			return nil, errors.Wrap(optionsErr, "could not set options for webp conversion")
		}

		imgBuffer := new(bytes.Buffer)
		if encodeErr := webp.Encode(imgBuffer, img, webpOptions); encodeErr != nil {
			return nil, errors.Wrap(encodeErr, "unable to encode the given image to webp")
		}

		return imgBuffer.Bytes(), nil
	case "image/jpeg":
		img, err := jpeg.Decode(bytes.NewReader(image))
		if err != nil {
			return nil, errors.Wrap(err, "unable to decode jpeg")
		}

		webpOptions, optionsErr := encoder.NewLosslessEncoderOptions(encoder.PresetDefault, 75)
		if optionsErr != nil {
			return nil, errors.Wrap(optionsErr, "could not set options for webp conversion")
		}

		imgBuffer := new(bytes.Buffer)
		if encodeErr := webp.Encode(imgBuffer, img, webpOptions); encodeErr != nil {
			return nil, errors.Wrap(encodeErr, "unable to encode the given image to webp")
		}

		return imgBuffer.Bytes(), nil
	}

	return nil, fmt.Errorf("unable to convert %#v to png", imgType)
}
