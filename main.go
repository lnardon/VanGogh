package main

import (
	"fmt"
	"image"
	"os"
	"strings"

	"golang.org/x/image/draw"

	//Local imports
	"github.com/lnardon/VanGogh/decoder"
	"github.com/lnardon/VanGogh/encoder"
)

func main() {
	fmt.Println("VanGogh - An image manipulation tool written in Go.")
    args := os.Args

    input, _ := os.Open(args[1])
    defer input.Close()
    outputFile, _ := os.Create("image_resized.png")
    defer outputFile.Close()
    extension := input.Name()[strings.LastIndex(input.Name(),"."):]
    src,err := decoder.DecodeImage(input,extension)
    
    if err != nil {
      fmt.Println(err)
      } else {
      fmt.Println("Resizing...")
      blankImage := image.NewRGBA(image.Rect(0, 0, src.Bounds().Max.X/10, src.Bounds().Max.Y/10))
      draw.NearestNeighbor.Scale(blankImage, blankImage.Rect, src, src.Bounds(), draw.Over, nil)
      encoder.EncodeImage(outputFile,blankImage, extension)
      fmt.Println("Done.")
    }
}

