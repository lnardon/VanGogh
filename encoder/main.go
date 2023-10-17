package encoder

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func EncodeImage(outputFile *os.File, newImage image.Image, extension string)(error){
  switch extension {
    case ".png":
      png.Encode(outputFile, newImage)
      return nil
    case ".jpg":
      jpeg.Encode(outputFile,newImage,nil)
      return nil
    case ".jpeg":
      jpeg.Encode(outputFile, newImage,nil)
      return nil
    default:
    return nil
  } 
}
