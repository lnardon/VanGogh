package main

import (
	"fmt"
    "os"
    "image/png"
    "golang.org/x/image/draw"
    "image"
)

func main() {
	fmt.Println("VanGogh - An image manipulation tool written in Go.")
    args := os.Args

    input, _ := os.Open(args[1])
    defer input.Close()
    output, _ := os.Create("image_resized.png")
    defer output.Close()
    
    fmt.Println("Resizing...")
    src, _ := png.Decode(input)
    dst := image.NewRGBA(image.Rect(0, 0, src.Bounds().Max.X/2, src.Bounds().Max.Y/2))
    draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)      
    png.Encode(output, dst)
	fmt.Println("Done.")
}
