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
	
	fileInfo, err := os.Stat(os.Args[1])
	isBatchOperation := err == nil && fileInfo.IsDir()

	if isBatchOperation {
		// TODO: Make it work with goroutines(parallel)
		fmt.Println("Starting batch operation...")
		args := os.Args
		scalingTechnique := args[2]
		scalingFactor, _ := strconv.Atoi(args[3])
		files,_:= os.ReadDir((args[1]))

		for _, file := range files {
			filePath := args[1] + "/" + file.Name()
			fileName := file.Name()
			extension := fileName[strings.LastIndex(fileName, "."):]
			input, _ := os.Open(filePath)
			defer input.Close()
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			
			outputFile, _ := os.Create("output/" + fileName)
			defer outputFile.Close()
			
			src, err := decoder.DecodeImage(input, extension)
			if err != nil {
				fmt.Println(err,"Error decoding image.")
			}
			
			newImage := scaler.ScaleImage(src, scalingTechnique, scalingFactor)
			encoder.EncodeImage(outputFile, newImage, extension)
		}
		fmt.Println("Finished!")

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

		newImage := scaler.ScaleImage(src, scalingTechnique, scalingFactor)
		encoder.EncodeImage(outputFile, newImage, extension)
		fmt.Println("Finished!")
	}
}
