package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"golang.org/x/image/draw"
)

func decodeImage(input *os.File, extension string) (image.Image, error) {
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

func encodeImage(outputFile *os.File, newImage image.Image, extension string)(error){
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

func main() {
	fmt.Println("VanGogh - An image manipulation tool written in Go.")
    args := os.Args

    input, _ := os.Open(args[1])
    defer input.Close()
    outputFile, _ := os.Create("image_resized.png")
    defer outputFile.Close()
    extension := args[2]
    
    fmt.Println("Resizing...")
    src,err := decodeImage(input,extension)
    if err != nil {
        fmt.Println(err)
    } else {
        blankImage := image.NewRGBA(image.Rect(0, 0, src.Bounds().Max.X/2, src.Bounds().Max.Y/2))
        draw.NearestNeighbor.Scale(blankImage, blankImage.Rect, src, src.Bounds(), draw.Over, nil)
        encodeImage(outputFile,blankImage, extension)
        fmt.Println("Done.")
    }
}

