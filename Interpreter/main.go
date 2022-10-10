package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Element interface {
	Value() int
}

type Integer struct {
	value int
}

func (i *Integer) Value() int {
	return i.value
}

type Operation int

const (
	Addition Operation = iota
	Subtraction
)

type BinaryOperation struct {
	left, right Element
	op          Operation
}

func (b *BinaryOperation) Value() int {
	switch b.op {
	case Addition:
		return b.left.Value() + b.right.Value()
	case Subtraction:
		return b.left.Value() - b.right.Value()
	default:
		panic("unknown operation")
	}
}

type TokenType int

const (
	Int TokenType = iota
	Plus
	Minus
	Lparen
	Rparen
)

type Token struct {
	Type TokenType
	Text string
}

func (t *Token) String() string {
	return fmt.Sprintf("'%s'", t.Text)
}

func lex(input string) []Token {
	var result []Token
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '+':
			result = append(result, Token{Plus, "+"})
		case '-':
			result = append(result, Token{Minus, "-"})
		case '(':
			result = append(result, Token{Lparen, "("})
		case ')':
			result = append(result, Token{Rparen, ")"})
		default:
			sb := strings.Builder{}
			for j := i; j < len(input); j++ {
				if unicode.IsDigit(rune(input[j])) {
					sb.WriteRune(rune(input[j]))
					i++
				} else {
					result = append(result, Token{Int, sb.String()})
					i--
					break
				}
			}
		}
	}
	return result
}

func parse(tokens []Token) Element {
	result := &BinaryOperation{
		left: &Integer{0},
		op:   Addition,
	}
	haveLHS := false
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		switch t.Type {
		case Int:
			intValue := 0
			fmt.Sscan(t.Text, &intValue)
			integer := &Integer{intValue}
			if !haveLHS {
				result.left = integer
				haveLHS = true
			} else {
				result.right = integer
			}
		case Plus:
			result.op = Addition
		case Minus:
			result.op = Subtraction
		case Lparen:
			j := i
			for ; j < len(tokens); j++ {
				if tokens[j].Type == Rparen {
					break
				}
			}
			subexpression := tokens[i+1 : j]
			element := parse(subexpression)
			if !haveLHS {
				result.left = element
				haveLHS = true
			} else {
				result.right = element
			}
			i = j
		}
	}
	return result
}

func main() {
	input := "(13+(4-1))-(12+(1-1))"
	tokens := lex(input)
	for _, t := range tokens {
		fmt.Println(t.String())
	}
	result := parse(tokens)
	fmt.Println(result.Value())

}
