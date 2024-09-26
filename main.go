package main

import (
	"fmt"
	"os"
)

func main() {
	train()
}

func train() {
	trainImagesPaths := GetTrainImagesPaths()

	neuron := NewNeuron()

	for i := 0; i < TrainImagesCount; i++ {
		trainWithImage(&neuron, trainImagesPaths.sharps[i], SharpValue)
		trainWithImage(&neuron, trainImagesPaths.dollars[i], DollarValue)
	}
}

func trainWithImage(neuron *Neuron, filepath string, expectedResult int) {
	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	imageMatrix := ConvertImageToMatrix(file)
	_ = file.Close()

	result := neuron.Consume(imageMatrix)

	if result != expectedResult {
		fmt.Printf("Error: %s\n", filepath)

		delta := expectedResult - result

		for i := 0; i < len(neuron.weights); i++ {
			neuron.weights[i] += LearningSpeed * float64(delta) * float64(imageMatrix[i/ImageSideSize][i%ImageSideSize])
		}
	} else {
		fmt.Printf("Predicted: %s\n", filepath)
	}
}
