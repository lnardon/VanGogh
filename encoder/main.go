package encoder

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/nickalie/go-webpbin"
)

func EncodeImage(outputFile *os.File, newImage image.Image, extension string)(error){
  switch extension {
    case ".png":
      png.Encode(outputFile, newImage)
      return nil
    case ".jpg":
      jpeg.Encode(outputFile,newImage, &jpeg.Options{Quality: 100})
      return nil
    case ".jpeg":
      jpeg.Encode(outputFile, newImage,&jpeg.Options{Quality: 100})
      return nil
    case ".webp":
      	if err := webpbin.Encode(outputFile, newImage); err != nil {
          outputFile.Close()
        }
      return nil
    default:
    return nil
  } 
}
