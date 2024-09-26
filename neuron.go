package main

import (
	bytes2 "bytes"
	"encoding/binary"
	"math/rand"
	"os"
)

type Neuron struct {
	weights [NeuronsQuantity]float64
}

func NewNeuron() Neuron {
	neuron := Neuron{weights: [NeuronsQuantity]float64{}}

	for i := 0; i < NeuronsQuantity; i++ {
		neuron.weights[i] = rand.Float64() - 0.5
	}

	return neuron
}

func (n Neuron) Consume(matrix [ImageSideSize][ImageSideSize]int8) int {
	sum := 0.0
	for i := 0; i < ImageSideSize; i++ {
		for j := 0; j < ImageSideSize; j++ {
			isBlack := matrix[i][j]

			if isBlack == 1 {
				sum += n.weights[i*ImageSideSize+j]
			}
		}
	}

	return activate(sum)
}

func (n Neuron) Save() {
	var bytes bytes2.Buffer
	if err := binary.Write(&bytes, binary.BigEndian, n.weights); err != nil {
		panic(err)
	}

	if err := os.WriteFile(WeightsFilename, bytes.Bytes(), 0666); err != nil {
		panic(err)
	}
}

func (n Neuron) Load() {
	f, err := os.Open(WeightsFilename)

	if err != nil {
		panic(err)
	}

	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	if err := binary.Read(f, binary.LittleEndian, &n.weights); err != nil {
		panic(err)
	}
}

func activate(value float64) int {
	if value > 0.0 {
		return 1
	} else {
		return 0
	}
}
