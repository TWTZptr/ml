package main

import (
	"fmt"
	"os"
)

func main() {
	neuron := NewNeuron()

	if !neuron.Load() {
		train(&neuron)
		fmt.Println("=================endof train=================")
	}

	run(&neuron)
}

func run(neuron *Neuron) {
	testImagesPaths := GetTestImagesPaths()

	for _, path := range testImagesPaths {
		result := predict(neuron, path)

		if result == SharpValue {
			fmt.Printf("%s is a Sharp\n", path)
		} else {
			fmt.Printf("%s is a Dollar\n", path)
		}
	}
}

func predict(neuron *Neuron, filepath string) int8 {
	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	imageMatrix := ConvertImageToMatrix(file)
	_ = file.Close()
	return neuron.Consume(imageMatrix)
}

func train(neuron *Neuron) {
	trainImagesPaths := GetTrainImagesPaths()

	for i := 0; i < TrainImagesCount; i++ {
		trainWithImage(neuron, trainImagesPaths.sharps[i], SharpValue)
		trainWithImage(neuron, trainImagesPaths.dollars[i], DollarValue)
	}

	neuron.Save()
}

func trainWithImage(neuron *Neuron, filepath string, expectedResult int8) {
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
