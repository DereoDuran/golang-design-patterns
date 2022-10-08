package main

import "fmt"

func main() {
	p := NewPersonBuilder()
	p.Lives().At("123 London Road").In("London").WithPostcode("SW12BC")
	p.Works().At("Fabrikam").AsA("Engineer").Earning(123000)
	person := p.Build()
	fmt.Println(person)
}
