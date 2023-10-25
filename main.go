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

	action := os.Args[1]

	if action == "convert" {
			fileInfo, err := os.Stat(os.Args[2])
			isBatchOperation := err == nil && fileInfo.IsDir()

			if isBatchOperation {
				fmt.Println("Starting batch operation...")
				args := os.Args
				targetExtension := args[3]
				files,_:= os.ReadDir((args[2]))

				for _, file := range files {
					filePath := args[2] + "/" + file.Name()
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
					
					newImage := scaler.ScaleImage(src, "BiLinear", 1)
					encoder.EncodeImage(outputFile, newImage, targetExtension)
				}
				fmt.Println("Finished!")
			} else {
				args := os.Args
				input, _ := os.Open(args[2])
				targetExtension := args[3]
				fmt.Println(targetExtension)
				fileNameWithExt := input.Name()[strings.LastIndex(input.Name(), "/")+1:]
				fileName := fileNameWithExt[:strings.LastIndex(fileNameWithExt, ".")]
				extension := fileNameWithExt[strings.LastIndex(fileNameWithExt, "."):]
				
				defer input.Close()
				outputFile, _ := os.Create("output/" + fileName + targetExtension)
				defer outputFile.Close()
				
				src, err := decoder.DecodeImage(input, extension)
				if err != nil {
					fmt.Println(err,"Error decoding image.")
				}
		
				newImage := scaler.ScaleImage(src, "BiLinear", 1)
				encoder.EncodeImage(outputFile, newImage, targetExtension)
				fmt.Println("Finished!")
			}
	} else {
	
	fileInfo, err := os.Stat(os.Args[2])
	isBatchOperation := err == nil && fileInfo.IsDir()

	if isBatchOperation {
		// TODO: Make it work with goroutines(parallel)
		fmt.Println("Starting batch operation...")
		args := os.Args
		scalingTechnique := args[3]
		scalingFactor, _ := strconv.Atoi(args[4])
		files,_:= os.ReadDir((args[2]))

		for _, file := range files {
			filePath := args[2] + "/" + file.Name()
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
		input, _ := os.Open(args[2])
		scalingTechnique := args[3]
		scalingFactor, _ := strconv.Atoi(args[4])
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
}
