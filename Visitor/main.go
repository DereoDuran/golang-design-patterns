package main

import (
	"fmt"
	"strings"
)

type Expression interface {
	Print(sb *strings.Builder)
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Print(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%g", d.value))
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Print(sb *strings.Builder) {
	sb.WriteString("(")
	a.left.Print(sb)
	sb.WriteString("+")
	a.right.Print(sb)
	sb.WriteString(")")
}

func Print(e Expression, sb *strings.Builder) {
	if de, ok := e.(*DoubleExpression); ok {
		sb.WriteString(fmt.Sprintf("%g", de.value))
	} else if ae, ok := e.(*AdditionExpression); ok {
		sb.WriteString("(")
		Print(ae.left, sb)
		sb.WriteString("+")
		Print(ae.right, sb)
		sb.WriteString(")")
	}
}

func main() {
	// 1 + (2+3)
	e := &AdditionExpression{
		left:  &DoubleExpression{1},
		right: &AdditionExpression{&DoubleExpression{2}, &DoubleExpression{3}},
	}
	sb := strings.Builder{}
	e.Print(&sb)
	fmt.Println(sb.String())

	sb.Reset()

	Print(e, &sb)
	fmt.Println(sb.String())
}
