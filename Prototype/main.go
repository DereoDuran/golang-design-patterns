package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = &Address{p.Address.StreetAddress, p.Address.City, p.Address.Country}
	return &q
}

func (p *Person) DeepCopySerialize() *Person {
	b := bytes.Buffer{}
	enc := gob.NewEncoder(&b)
	_ = enc.Encode(p)

	dec := gob.NewDecoder(&b)
	q := Person{}
	_ = dec.Decode(&q)
	return &q
}

type Address2 struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Address2
}

func (p *Employee) DeepCopySerialize() *Employee {
	b := bytes.Buffer{}
	enc := gob.NewEncoder(&b)
	_ = enc.Encode(p)

	dec := gob.NewDecoder(&b)
	q := Employee{}
	_ = dec.Decode(&q)
	return &q
}

var mainOffice = Employee{
	"", Address2{0, "123 East Dr", "London"},
}

var auxOffice = Employee{
	"", Address2{0, "66 West Dr", "London"},
}

func NewEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopySerialize()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return NewEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return NewEmployee(&auxOffice, name, suite)
}

func main() {
	john := Person{"John", &Address{"123 London Road", "London", "UK"}}
	jane := john
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker Street"

	fmt.Println(john, john.Address.StreetAddress)
	fmt.Println(jane, jane.Address.StreetAddress)

	jim := john.DeepCopy()
	jim.Name = "Jim"
	jim.Address.StreetAddress = "111 New York Avenue"

	fmt.Println(john, john.Address.StreetAddress)
	fmt.Println(jim, jim.Address.StreetAddress)

	jill := john.DeepCopySerialize()
	jill.Name = "Jill"
	jill.Address.StreetAddress = "222 San Francisco Street"

	fmt.Println(john, john.Address.StreetAddress)
	fmt.Println(jill, jill.Address.StreetAddress)

	johnAux := NewMainOfficeEmployee("John", 123)
	janeAux := NewAuxOfficeEmployee("Jane", 321)

	fmt.Println(johnAux)
	fmt.Println(janeAux)
}
