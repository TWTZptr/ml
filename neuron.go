package main

import (
	"bytes"
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

func (n *Neuron) Consume(matrix [ImageSideSize][ImageSideSize]int8) int8 {
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

func (n *Neuron) Save() {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, &n.weights); err != nil {
		panic(err)
	}

	if err := os.WriteFile(WeightsFilename, buf.Bytes(), 0666); err != nil {
		panic(err)
	}
}

func (n *Neuron) Load() bool {
	f, err := os.Open(WeightsFilename)

	if err != nil {
		return false
	}

	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	if err := binary.Read(f, binary.LittleEndian, &n.weights); err != nil {
		return false
	}

	return true
}

func activate(value float64) int8 {
	if value > 0.0 {
		return 1
	} else {
		return 0
	}
}
