package decoder

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func DecodeImage(input *os.File, extension string) (image.Image, error) {
	switch extension {
		case ".png":
			src, _ := png.Decode(input)
			return src,nil
		case ".jpg":
			src,_ := jpeg.Decode(input)
			return src,nil
		case ".jpeg":
			src,_ := jpeg.Decode(input)
			return src,nil
		default:
			return nil,errors.New("Unsuported file format!")
	}
}
