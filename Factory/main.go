package main

import "fmt"

type Person1 struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson1(name string, age int) *Person1 {
	return &Person1{Name: name, Age: age, EyeCount: 2}
}

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Println("Hello, my name is", p.name, "and I am", p.age, "years old.")
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func (p *tiredPerson) SayHello() {
	fmt.Println("Sorry, I am too tired.")
}

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

type EmployeeFactory struct {
	position     string
	annualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.position, f.annualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func main() {
	p := NewPerson("John Doe", 142)
	p.SayHello()

	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 80000)

	developer := developerFactory("Adam Smith")
	manager := managerFactory("Jane Doe")

	fmt.Println(developer)
	fmt.Println(manager)

	bossFactory := NewEmployeeFactory2("CEO", 100000)
	bossFactory.annualIncome = 200000
	boss := bossFactory.Create("Samuel Jackson")

	fmt.Println(boss)
}
