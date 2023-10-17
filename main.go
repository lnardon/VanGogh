package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	//Local imports
	"github.com/lnardon/VanGogh/decoder"
	"github.com/lnardon/VanGogh/encoder"
	"github.com/lnardon/VanGogh/scaler"
)

func main() {
	fmt.Println("VanGogh - An image manipulation tool written in Go.")
	args := os.Args

	input, _ := os.Open(args[1])
	scalingTechnique := args[2]
	scalingFactor, _ := strconv.Atoi(args[3])
	
	defer input.Close()
	outputFile, _ := os.Create("image_resized.png")
	defer outputFile.Close()
	extension := input.Name()[strings.LastIndex(input.Name(), "."):]
	src, err := decoder.DecodeImage(input, extension)

	if err != nil {
		fmt.Println(err)
	}

	info, err := os.Stat(args[1])
	isBathOperation := err == nil && info.IsDir()
	if isBathOperation {
		// TODO: HANDLE BATCH
		fmt.Println("FOI")
	} else {
		fmt.Println("Resizing...")
		newImage := scaler.ScaleImage(src, scalingTechnique, scalingFactor)
		encoder.EncodeImage(outputFile, newImage, extension)
		fmt.Println("Done.")
	}
}
