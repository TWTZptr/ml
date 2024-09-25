package main

type Perceptron struct {
	weight float64
}

func NewPerceptron() Perceptron {
	return Perceptron{weight: 1}
}
