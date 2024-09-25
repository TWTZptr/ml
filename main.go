package main

import "os"

func main() {
	trainImagesPaths := GetTrainImagesPaths()

	for _, path := range trainImagesPaths {
		file, err := os.Open(path)

		if err != nil {
			panic(err)
		}

		ConvertImageToMatrix(file)
		_ = file.Close()
	}
}
