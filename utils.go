package main

import (
	"image"
	_ "image/png"
	"os"
)

type TrainImages struct {
	sharps  []string
	dollars []string
}

func getFilePaths(path string) []string {
	entries, _ := os.ReadDir(path)
	res := make([]string, len(entries))
	for i, value := range entries {
		res[i] = path + "/" + value.Name()
	}

	return res
}

func GetTestImagesPaths() []string {
	return getFilePaths("images/test")
}

func GetTrainImagesPaths() TrainImages {
	return TrainImages{dollars: getFilePaths("images/train/dollar"), sharps: getFilePaths("images/train/sharp")}
}

func ConvertImageToMatrix(file *os.File) [ImageSideSize][ImageSideSize]int8 {
	imageData, _, err := image.Decode(file)

	if err != nil {
		panic(err)
	}

	stepX, stepY := float64(imageData.Bounds().Max.X)/ImageSideSize, float64(imageData.Bounds().Max.Y)/ImageSideSize

	result := [ImageSideSize][ImageSideSize]int8{}

	for i := 0.0; i < ImageSideSize; i++ {
		for j := 0.0; j < ImageSideSize; j++ {
			pixel := imageData.At(int(i*stepX), int(j*stepY))
			_, r, g, b := pixel.RGBA()
			sum := r + g + b
			if sum < ColorThreshold {
				result[int(i)][int(j)] = 1
			} else {
				result[int(i)][int(j)] = 0
			}
		}
	}

	return result
}
