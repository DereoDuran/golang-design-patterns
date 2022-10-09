package main

import (
	"fmt"
	"strings"
)

type GraphicObject struct {
	Name, Color string
	Children    []GraphicObject
}

func (g *GraphicObject) Print(indent int) {
	sb := strings.Builder{}
	for i := 0; i < indent; i++ {
		sb.WriteString(" ")
	}
	sb.WriteString(g.Name)
	if g.Color != "" {
		sb.WriteString(" (")
		sb.WriteString(g.Color)
		sb.WriteString(")")
	}
	fmt.Println(sb.String())
	for _, c := range g.Children {
		c.Print(indent + 2)
	}
}

func NewCircle(color string) *GraphicObject {
	return &GraphicObject{"Circle", color, nil}
}

func NewSquare(color string) *GraphicObject {
	return &GraphicObject{"Square", color, nil}
}

type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}

type NeuronLayer struct {
	Neurons []Neuron
}

func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{make([]Neuron, count)}
}

type NeuronInterface interface {
	Iter() []*Neuron
}

func (n *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, len(n.Neurons))
	for i := range n.Neurons {
		result[i] = &n.Neurons[i]
	}
	return result
}

func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

func ConnectToAll(from, to NeuronInterface) {
	for _, source := range from.Iter() {
		for _, dest := range to.Iter() {
			source.ConnectTo(dest)
		}
	}
}

func main() {
	drawing := GraphicObject{"My Drawing", "", nil}
	drawing.Children = append(drawing.Children, *NewSquare("Red"))

	group := GraphicObject{"Group 1", "", nil}
	group.Children = append(group.Children, *NewCircle("Blue"))
	group.Children = append(group.Children, *NewSquare("Blue"))

	drawing.Children = append(drawing.Children, group)

	drawing.Children = append(drawing.Children, *NewCircle("Yellow"))

	drawing.Print(0)

	neuron1, neuron2 := Neuron{}, Neuron{}
	layer1, layer2 := NewNeuronLayer(3), NewNeuronLayer(4)

	neuron1.ConnectTo(&neuron2)
	ConnectToAll(layer1, &neuron1)
	ConnectToAll(&neuron2, layer2)
	ConnectToAll(layer1, layer2)

	fmt.Println("Neuron 1:", len(neuron1.In), len(neuron1.Out))
	fmt.Println("Neuron 2:", len(neuron2.In), len(neuron2.Out))
	fmt.Println("Layer 1:", len(layer1.Neurons))
	fmt.Println("Layer 2:", len(layer2.Neurons))

}
