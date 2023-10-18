package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lnardon/VanGogh/decoder"
	"github.com/lnardon/VanGogh/encoder"
	"github.com/lnardon/VanGogh/scaler"
)

func main() {
	fmt.Println("VanGogh - An image manipulation tool written in Go.")
	
	info, err := os.Stat(os.Args[1])
	isBatchOperation := err == nil && info.IsDir()
	if isBatchOperation {
		// TODO: handle batch and make it work with goroutines
		fmt.Println("Unable to handle batch operations yet.")
	} else {
	  args := os.Args
      input, _ := os.Open(args[1])
      scalingTechnique := args[2]
      scalingFactor, _ := strconv.Atoi(args[3])
      fileName := input.Name()[strings.LastIndex(input.Name(), "/")+1:]
      extension := fileName[strings.LastIndex(fileName, "."):]
      
      defer input.Close()
      outputFile, _ := os.Create("output/" + fileName)
      defer outputFile.Close()
      
	  src, err := decoder.DecodeImage(input, extension)
      if err != nil {
        fmt.Println(err,"Error decoding image.")
      }

      fmt.Println("Resizing...")
      newImage := scaler.ScaleImage(src, scalingTechnique, scalingFactor)
      encoder.EncodeImage(outputFile, newImage, extension)
      fmt.Println("Done.")
	}
}
