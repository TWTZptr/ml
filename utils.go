package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
)

func getFilesPaths(path string) []string {
	entries, _ := os.ReadDir(path)
	res := make([]string, len(entries))
	for i, value := range entries {
		res[i] = path + "/" + value.Name()
	}

	return res
}

func GetTestImagesPaths() []string {
	return getFilesPaths("images/test")
}

func GetTrainImagesPaths() []string {
	return getFilesPaths("images/train")
}

func ConvertImageToMatrix(file *os.File) {
	imageData, imageType, err := image.Decode(file)

	if err != nil {
		panic(err)
	}

	fmt.Println(imageType, imageData)
}
